package handler

import (
	"context"
	"strconv"

	dbUser "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/user"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/common"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create new user
func (h *Handler) CreateUser(ctx context.Context, in *user.User) (userId *user.Id, err error) {

	innerUser := &dbUser.UserFullStuff{}
	_ = innerUser.ProtoToDb(in)
	transit := &dbUser.TransitUser{
		U:                  innerUser,
		OriginalPhoto:      in.GetPhoto(),
		OriginalResumeFile: in.Resume.GetUploadedResume()}
	innerId, err := h.services.Mnemosyne.AddUser(ctx, transit)
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

	ur.WithContacts = in.Option.GetWithContacts()
	ur.WithResume = in.Option.GetWithResume()
	ur.WithDeleted = in.Option.GetWithDeleted()

	ur.Role = in.Role.GetRole()

	ur.FieldName = in.Filter.GetFieldName()
	ur.FieldValue = in.Filter.GetFieldValue()

	ur.Group = in.Option.GetGroupId()
	ur.Team = in.Option.GetTeamId()
	innerUsers, err := h.services.Mnemosyne.GetUsers(ctx, ur)
	if len(innerUsers) > 0 {
		for _, user := range innerUsers {
			innerApiUser := user.DbToProto()
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
func (h *Handler) GetUserById(ctx context.Context, in *user.Id) (u *user.User, err error) {
	u = new(user.User)
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		return u, status.Error(codes.InvalidArgument, err.Error())
	}
	innerUser, err := h.services.Mnemosyne.GetUserById(ctx, innerId)
	if innerUser != nil {
		u = innerUser.DbToProto()
	}
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return
	}
	return
}

// Get user by email
func (h *Handler) GetUserByEmail(ctx context.Context, in *user.Email) (u *user.User, err error) {
	u = &user.User{}
	email := in.Email
	innerUser, err := h.services.Mnemosyne.GetUserByEmail(ctx, email)
	if innerUser != nil {
		u = innerUser.DbToProto()
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
	transit := &dbUser.TransitUser{
		U:                  innerUser,
		OriginalPhoto:      in.GetPhoto(),
		OriginalResumeFile: in.Resume.GetUploadedResume()}

	err = h.services.Mnemosyne.UpdateUser(ctx, transit)
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
	c = &user.Contact{}
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	innerContact, err := h.services.Mnemosyne.GetContactById(ctx, innerId)
	if innerContact != nil {
		c = &user.Contact{
			Id:                   strconv.Itoa(innerContact.Id),
			Telegram:             innerContact.Telegram,
			Discord:              innerContact.Discord,
			CommunicationChannel: innerContact.CommunicationChannel,
		}
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
		Telegram:             in.GetTelegram(),
		Discord:              in.GetDiscord(),
		CommunicationChannel: in.GetCommunicationChannel(),
	})

	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Get resume by ID
func (h *Handler) GetResume(ctx context.Context, in *user.Id) (r *user.Resume, err error) {
	r = &user.Resume{}
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	innerResume, err := h.services.Mnemosyne.GetResumeById(ctx, innerId)
	if innerResume != nil {
		r = &user.Resume{
			Id:             strconv.Itoa(innerResume.Id),
			UploadedResume: &common.File{Name: innerResume.UploadedResume},
			Experience:     innerResume.Experience,
			Country:        innerResume.Country,
			City:           innerResume.City,
			TimeZone:       innerResume.TimeZone,
			MentorsNote:    innerResume.MentorsNote,
		}
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return
}

// Update resume data
func (h *Handler) UpdateResume(ctx context.Context, in *user.Resume) (c *common.Empty, err error) {
	c = &common.Empty{}
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}

	err = h.services.Mnemosyne.UpdateResume(ctx, &dbUser.TransitResume{
		OriginalResumeFile: in.GetUploadedResume(),
		R: &dbUser.Resume{
			Id:          innerId,
			Experience:  in.GetExperience(),
			Country:     in.GetExperience(),
			City:        in.GetCity(),
			TimeZone:    in.GetTimeZone(),
			MentorsNote: in.GetMentorsNote(),
		},
	})

	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Delete contact by ID
func (h *Handler) DeleteContact(ctx context.Context, in *user.Id) (c *common.Empty, err error) {
	c = &common.Empty{}
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = h.services.Mnemosyne.DeleteContact(ctx, innerId)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}

// Delete resume by ID
func (h *Handler) DeleteResume(ctx context.Context, in *user.Id) (c *common.Empty, err error) {
	c = &common.Empty{}
	innerId, err := strconv.Atoi(in.Id)
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = h.services.Mnemosyne.DeleteResume(ctx, innerId)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return
}
