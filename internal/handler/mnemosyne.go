package handler

import (
	"context"
	"strconv"

	"log"

	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
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

	dbUser, _ := dbUser.ProtoUserToDbUser(in)
	id, err := h.services.Mnemosyne.AddUser(ctx, *dbUser)
	userId = &(user.Id{Id: strconv.Itoa(id)})
	return
}

// Get all existing users
func (h *Handler) GetUsers(ctx context.Context, in *user.UserRequest) (users *user.Users, err error) {

	innerUserrequest := &(dbUser.UserRequest{WithContacts: in.Option.WithContacts,
		WithResume:  in.Option.WithResume,
		WithDeleted: in.Option.WithDeleted,
		Role:        in.Role.Role,
		FieldName:   in.Filter.FieldName,
		FieldValue:  in.Filter.FieldValue})
	dbUsers, err := h.services.Mnemosyne.GetUsers(ctx, innerUserrequest)
	var structUsers []*user.User
	for _, value := range dbUsers {
		tempUser := value
		structUser := dbUser.DbUserToProtoUser(&tempUser)
		structUsers = append(structUsers, structUser)
	}
	users = &user.Users{Users: structUsers}
	return
}

// Get user by id
func (h *Handler) GetUserById(ctx context.Context, in *user.Id) (user *user.User, err error) {
	innerId, innerErr := strconv.Atoi(in.Id)
	if innerErr != nil {
		return nil, errors.Wrap(innerErr, "invalid user's id value")
	}
	innerUser, err := h.services.Mnemosyne.GetUserById(ctx, innerId)
	user = dbUser.DbUserToProtoUser(&innerUser)
	return
}

// Get user by email
func (h *Handler) GetUserByEmail(ctx context.Context, in *user.Email) (user *user.User, err error) {
	innerUser, err := h.services.Mnemosyne.GetUserByEmail(ctx, in.Email)
	user = dbUser.DbUserToProtoUser(&innerUser)
	return
}

// Update user's data
func (h *Handler) UpdateUser(ctx context.Context, in *user.User) (ok *wrapperspb.BoolValue, err error) {
	innerUser, err := dbUser.ProtoUserToDbUser(in)
	if err != nil {
		err = errors.Wrap(err, "crushed on parsing user's info")
		ok = &wrapperspb.BoolValue{Value: false}
		return
	}
	innerOk, err := h.services.Mnemosyne.UpdateUser(ctx, *innerUser)
	ok = &wrapperspb.BoolValue{Value: innerOk}
	return
}

// Delete user by id
func (h *Handler) DeleteUser(ctx context.Context, in *user.Id) (ok *wrapperspb.BoolValue, err error) {
	innerId, innerErr := strconv.Atoi(in.Id)
	if innerErr != nil {
		err = errors.Wrap(innerErr, "invalid user's id value")
		ok = &wrapperspb.BoolValue{Value: false}
		return
	}
	innerOk, err := h.services.Mnemosyne.DeleteUser(ctx, innerId)
	ok = &wrapperspb.BoolValue{Value: innerOk}
	return
}

// Delete user by id
func (h *Handler) ActivateUser(ctx context.Context, in *user.Id) (ok *wrapperspb.BoolValue, err error) {
	innerId, innerErr := strconv.Atoi(in.Id)
	if innerErr != nil {
		err = errors.Wrap(innerErr, "invalid user's id value")
		ok = &wrapperspb.BoolValue{Value: false}
		return
	}
	innerOk, err := h.services.Mnemosyne.ActivateUser(ctx, innerId)
	ok = &wrapperspb.BoolValue{Value: innerOk}
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
