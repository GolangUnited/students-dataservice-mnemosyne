package handler

import (
	"context"
	"strconv"
	"strings"

	"log"

	db "github.com/NEKETSKY/mnemosyne/models/database"
	"github.com/NEKETSKY/mnemosyne/models/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/pkg/api/helloworld"
	"github.com/NEKETSKY/mnemosyne/pkg/api/user"
	"github.com/NEKETSKY/mnemosyne/pkg/auth"
	"github.com/NEKETSKY/mnemosyne/pkg/operations"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// SayHello implements api.MnemosyneServer
func (h *Handler) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	user := auth.GetUser(ctx)
	_ = user

	access := operations.CheckAccess(ctx, "view_all_students")
	_ = access

	var req mnemosyne.Request
	resp, err := h.services.Mnemosyne.Test(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "say hello")
	}

	log.Printf("Received: %v", in.GetName())
	log.Printf("Version: %v", resp.Version)
	return &helloworld.HelloReply{Message: "Hello " + in.GetName() + ". Version " + resp.Version}, nil
}

// Create new user
func (h *Handler) CreateUser(ctx context.Context, in *user.User) (userId *user.Id, err error) {
	dbUser, _ := protoUserToDbUser(in)

	id, err := h.services.Mnemosyne.AddUser(ctx, *dbUser)

	userId = &(user.Id{Id: strconv.Itoa(id)})
	return
}

// Get all existing users
func (h *Handler) GetUsers(ctx context.Context, in *user.UserRequest) (users *user.Users, err error) {

	dbUsers, err := h.services.Mnemosyne.GetUsers(ctx)
	var structUsers []*user.User
	for _, value := range dbUsers {
		structUser := dbUserToProtoUser(&value)
		structUsers = append(structUsers, structUser)
	}
	users.Users = structUsers
	return
}

// Get user by id
func (h *Handler) GetUserById(ctx context.Context, in *user.Id) (user *user.User, err error) {
	innerId, innerErr := strconv.Atoi(in.Id)
	if innerErr != nil {
		return nil, errors.Wrap(innerErr, "invalid user's id value")
	}
	innerUser, err := h.services.Mnemosyne.GetUserById(ctx, innerId)
	user = dbUserToProtoUser(&innerUser)
	return
}

// Get user by email
func (h *Handler) GetUserByEmail(ctx context.Context, in *user.Email) (user *user.User, err error) {
	innerUser, err := h.services.Mnemosyne.GetUserByEmail(ctx, in.Email)
	user = dbUserToProtoUser(&innerUser)
	return
}

// Update user's data
func (h *Handler) UpdateUser(ctx context.Context, in *user.User) (ok *wrapperspb.BoolValue, err error) {
	innerUser, err := protoUserToDbUser(in)
	if err != nil {
		ok.Value = false
		return
	}
	ok.Value, err = h.services.Mnemosyne.UpdateUser(ctx, *innerUser)
	return
}

// Delete user by id
func (h *Handler) DeleteUser(ctx context.Context, in *user.Id) (ok *wrapperspb.BoolValue, err error) {
	innerId, innerErr := strconv.Atoi(in.Id)
	if innerErr != nil {
		err = errors.Wrap(innerErr, "invalid user's id value")
		ok.Value = false
		return
	}
	ok.Value, err = h.services.Mnemosyne.DeleteUser(ctx, innerId)
	return
}

// Get contact by ID
func (h *Handler) GetContact(ctx context.Context, in *user.Id) (contact *user.Contact, err error) {
	return
}

// Update contact's data
func (h *Handler) UpdateContact(ctx context.Context, in *user.Contact) (ok *wrapperspb.BoolValue, err error) {
	return
}

func dbUserToProtoUser(d *db.User) (u *user.User) {
	u = &(user.User{Id: strconv.Itoa(d.Id),
		LastName:     d.LastName,
		FirstName:    d.FirstName,
		MiddleName:   &d.MiddleName,
		Email:        d.Email,
		Language:     d.Language,
		EnglishLevel: d.EnglishLevel,
		Photo:        d.Photo,
	})
	return
}
func protoUserToDbUser(u *user.User) (d *db.User, err error) {
	var innerId int
	innerId, err = strconv.Atoi(u.Id)
	if err != nil {
		err = errors.Wrap(err, "invalid user's id value")
	}
	var engLevel strings.Builder
	for index, value := range u.EnglishLevel {
		if index > 2 {
			break
		}
		engLevel.WriteRune(value)
	}
	d = &db.User{
		Id:           innerId,
		LastName:     u.LastName,
		FirstName:    u.FirstName,
		MiddleName:   *u.MiddleName,
		Email:        u.Email,
		Language:     u.Language,
		EnglishLevel: engLevel.String(),
		Photo:        u.Photo,
	}
	return
}
