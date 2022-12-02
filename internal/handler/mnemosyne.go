package handler

import (
	"context"
	"strconv"

	"log"

	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
	"github.com/NEKETSKY/mnemosyne/models/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/pkg/api/common"
	"github.com/NEKETSKY/mnemosyne/pkg/api/helloworld"
	"github.com/NEKETSKY/mnemosyne/pkg/api/user"
	"github.com/NEKETSKY/mnemosyne/pkg/auth"
	"github.com/NEKETSKY/mnemosyne/pkg/file"
	"github.com/NEKETSKY/mnemosyne/pkg/operations"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	innerUser := &dbUser.UserFullStuff{}
	_ = innerUser.ProtoToDb(in)
	innerId, err := h.services.Mnemosyne.AddUser(ctx, innerUser)
	userId = &user.Id{Id: strconv.Itoa(innerId)}
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Get all existing users
func (h *Handler) GetUsers(ctx context.Context, in *user.UserRequest) (users *user.Users, err error) {
	ur := &dbUser.UserRequest{}
	var innerApiUserSlice []*user.User
	if in.Option != nil {
		ur.WithContacts = in.Option.WithContacts
		ur.WithResume = in.Option.WithResume
	}
	if in.Role != nil {
		ur.Role = in.Role.Role
	}
	if in.Filter != nil {
		ur.FieldName = in.Filter.FieldName
		ur.FieldValue = in.Filter.FieldValue
	}
	innerUsers, err := h.services.Mnemosyne.GetUsers(ctx, ur)
	if len(innerUsers) > 0 {
		for _, user := range innerUsers {
			temp := user
			innerApiUser := temp.DbToProto()
			if !ur.WithContacts {
				innerApiUser.Contact = nil
			}
			if !ur.WithResume {
				innerApiUser.Resume = nil
			}
			innerApiUserSlice = append(innerApiUserSlice, innerApiUser)
		}
	}
	users = &user.Users{Users: innerApiUserSlice}
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Get user by id
func (h *Handler) GetUserById(ctx context.Context, in *user.Id) (user *user.User, err error) {

	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	innerUser, err := h.services.Mnemosyne.GetUserById(ctx, innerId)
	if innerUser != nil {
		user = innerUser.DbToProto()
	}
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Get user by email
func (h *Handler) GetUserByEmail(ctx context.Context, in *user.Email) (user *user.User, err error) {
	email := in.Email
	innerUser, err := h.services.Mnemosyne.GetUserByEmail(ctx, email)
	if innerUser != nil {
		user = innerUser.DbToProto()
	}
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Update user's data
func (h *Handler) UpdateUser(ctx context.Context, in *user.User) (c *common.Empty, err error) {
	c = &common.Empty{}
	innerUser := &dbUser.UserFullStuff{}
	err = innerUser.ProtoToDb(in)
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = h.services.Mnemosyne.UpdateUser(ctx, innerUser)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Delete user by id
func (h *Handler) DeactivateUser(ctx context.Context, in *user.Id) (c *common.Empty, err error) {
	c = &common.Empty{}
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = h.services.Mnemosyne.DeactivateUser(ctx, innerId)

	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Delete user by id
func (h *Handler) ActivateUser(ctx context.Context, in *user.Id) (c *common.Empty, err error) {
	c = &common.Empty{}
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = h.services.Mnemosyne.ActivateUser(ctx, innerId)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Get contact by ID
func (h *Handler) GetContact(ctx context.Context, in *user.Id) (c *user.Contact, err error) {

	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	innerContact, err := h.services.Mnemosyne.GetContactById(ctx, innerId)
	c = &user.Contact{
		Id:                   strconv.Itoa(innerContact.Id),
		Telegram:             innerContact.Telegram,
		Discord:              innerContact.Discord,
		CommunicationChannel: innerContact.CommunicationChannel,
	}
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Update contact's data
func (h *Handler) UpdateContact(ctx context.Context, in *user.Contact) (c *common.Empty, err error) {

	c = &common.Empty{}
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}

	err = h.services.Mnemosyne.UpdateContact(ctx, &dbUser.Contact{
		Id:                   innerId,
		Telegram:             in.Telegram,
		Discord:              in.Discord,
		CommunicationChannel: in.CommunicationChannel,
	})

	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Get resume by ID
func (h *Handler) GetResume(ctx context.Context, in *user.Id) (r *user.Resume, err error) {

	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	innerResume, err := h.services.Mnemosyne.GetResumeById(ctx, innerId)
	r = &user.Resume{
		Id:             strconv.Itoa(innerResume.Id),
		UploadedResume: &common.File{Name: innerResume.UploadedResume},
		Experience:     innerResume.Experience,
		Country:        innerResume.Country,
		City:           innerResume.City,
		TimeZone:       innerResume.TimeZone,
		MentorsNote:    innerResume.MentorsNote,
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return
}

// Update resume data
func (h *Handler) UpdateResume(ctx context.Context, in *user.Resume) (c *common.Empty, err error) {
	c = &common.Empty{}
	path := ""
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	if in.UploadedResume != nil {
		path, _ = file.Save(in.UploadedResume.GetName(), in.UploadedResume.GetContent())
	}
	err = h.services.Mnemosyne.UpdateResume(ctx, &dbUser.Resume{
		Id:             innerId,
		UploadedResume: path,
		Experience:     in.GetExperience(),
		Country:        in.GetExperience(),
		City:           in.GetCity(),
		TimeZone:       in.GetTimeZone(),
		MentorsNote:    in.GetMentorsNote(),
	})

	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}
