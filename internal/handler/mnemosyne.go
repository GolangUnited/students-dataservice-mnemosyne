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
	var engLevel strings.Builder
	for index, value := range in.EnglishLevel {
		if index > 2 {
			break
		}
		engLevel.WriteRune(value)
	}

	id, err := h.services.Mnemosyne.AddUser(ctx, db.User{
		LastName:     in.LastName,
		FirstName:    in.FirstName,
		MiddleName:   *in.MiddleName,
		Email:        in.Email,
		Language:     in.Language,
		EnglishLevel: engLevel.String(),
		Photo:        in.Photo,
	})

	userId = &(user.Id{Id: strconv.Itoa(id)})
	return
}

// Get all existing users
func (h *Handler) GetUsers(ctx context.Context, in *user.Options) (users *user.Users, err error) {
	dbUsers, err := h.services.Mnemosyne.GetUsers(ctx)
	var structUsers []*user.User
	for _, value := range dbUsers {
		structUser := &(user.User{Id: strconv.Itoa(value.Id),
			LastName:     value.LastName,
			FirstName:    value.FirstName,
			MiddleName:   &value.MiddleName,
			Email:        value.Email,
			Language:     value.Language,
			EnglishLevel: value.EnglishLevel,
			Photo:        value.Photo,
		})
		structUsers = append(structUsers, structUser)
	}
	users = &user.Users{Users: structUsers}
	return
}

// Get user by id
func (h *Handler) GetUser(ctx context.Context, in *user.Id) (user *user.User, err error) {
	return
}

// Update user's data
func (h *Handler) UpdateUser(ctx context.Context, in *user.User) (ok *wrapperspb.BoolValue, err error) {
	return
}

// Delete user by id
func (h *Handler) DeleteUser(ctx context.Context, in *user.Id) (ok *wrapperspb.BoolValue, err error) {
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
