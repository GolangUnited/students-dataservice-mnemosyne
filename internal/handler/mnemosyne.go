package handler

import (
	"context"
	"strconv"

	"log"

	"github.com/NEKETSKY/mnemosyne/models/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/pkg/api/helloworld"
	"github.com/NEKETSKY/mnemosyne/pkg/api/user"
	"github.com/NEKETSKY/mnemosyne/pkg/auth"
	"github.com/NEKETSKY/mnemosyne/pkg/file"
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

	f, _ := file.Save("mytest.txt", []byte("mytest"))
	_ = f

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

	userId, err = h.services.Mnemosyne.AddUser(ctx, in)

	return
}

// Get all existing users
func (h *Handler) GetUsers(ctx context.Context, in *user.UserRequest) (users *user.Users, err error) {

	if access := operations.CheckAccess(ctx, "view_deleted"); !access {
		in.Option.WithDeleted = false
	}
	if access := operations.CheckAccess(ctx, "view_all_students"); access {
		users, err = h.services.Mnemosyne.GetUsers(ctx, in)
	} else {
		return nil, errors.New("access denied")
	}

	return
}

// Get user by id
func (h *Handler) GetUserById(ctx context.Context, in *user.Id) (user *user.User, err error) {

	check := auth.GetUser(ctx)

	innerId, innerErr := strconv.Atoi(in.Id)
	if innerErr != nil {
		return nil, errors.Wrap(innerErr, "invalid user's id value")
	}

	if (innerId != check.Id) && !(operations.CheckAccess(ctx, "view_all_students")) {
		return nil, errors.New("access denied")
	}

	user, err = h.services.Mnemosyne.GetUserById(ctx, innerId)
	return
}

// Get user by email
func (h *Handler) GetUserByEmail(ctx context.Context, in *user.Email) (user *user.User, err error) {

	check := auth.GetUser(ctx)
	checkEmail, _ := h.services.Mnemosyne.GetEmailById(ctx, check.Id)

	if (checkEmail.Email != in.Email) && !(operations.CheckAccess(ctx, "view_all_students")) {
		return nil, errors.New("access denied")
	}
	user, err = h.services.Mnemosyne.GetUserByEmail(ctx, in.Email)

	return
}

// Update user's data
func (h *Handler) UpdateUser(ctx context.Context, in *user.User) (ok *wrapperspb.BoolValue, err error) {

	check := auth.GetUser(ctx)

	innerId, innerErr := strconv.Atoi(in.Id)
	if innerErr != nil {
		return nil, errors.Wrap(innerErr, "invalid user's id value")
	}

	if (innerId != check.Id) && !(operations.CheckAccess(ctx, "view_all_students")) {
		return nil, errors.New("access denied")
	}

	innerOk, err := h.services.Mnemosyne.UpdateUser(ctx, in)
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
func (h *Handler) GetContact(ctx context.Context, in *user.Id) (c *user.Contact, err error) {
	check := auth.GetUser(ctx)
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid user's id value")
	}
	if (check.Id != innerId) && !(operations.CheckAccess(ctx, "view_all_students")) {
		return nil, errors.New("access denied")
	}
	c, err = h.services.Mnemosyne.GetContactById(ctx, innerId)
	return
}

// Update contact's data
func (h *Handler) UpdateContact(ctx context.Context, in *user.Contact) (ok *wrapperspb.BoolValue, err error) {
	check := auth.GetUser(ctx)
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid user's id value")
	}
	if (check.Id != innerId) && !(operations.CheckAccess(ctx, "view_all_students")) {
		return nil, errors.New("access denied")
	}
	innerOk, err := h.services.Mnemosyne.UpdateContact(ctx, in)
	ok = &wrapperspb.BoolValue{Value: innerOk}
	return
}

// Get resume by ID
func (h *Handler) GetResume(ctx context.Context, in *user.Id) (r *user.Resume, err error) {
	check := auth.GetUser(ctx)
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid user's id value")
	}
	if (check.Id != innerId) && !(operations.CheckAccess(ctx, "view_all_students")) {
		return nil, errors.New("access denied")
	}
	r, err = h.services.Mnemosyne.GetResumeById(ctx, innerId)
	return
}

// Update resume data
func (h *Handler) UpdateResume(ctx context.Context, in *user.Resume) (ok *wrapperspb.BoolValue, err error) {
	check := auth.GetUser(ctx)
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid user's id value")
	}
	if (check.Id != innerId) && !(operations.CheckAccess(ctx, "view_all_students")) {
		return nil, errors.New("access denied")
	}

	innerOk, err := h.services.Mnemosyne.UpdateResume(ctx, in)
	ok = &wrapperspb.BoolValue{Value: innerOk}
	return
}
