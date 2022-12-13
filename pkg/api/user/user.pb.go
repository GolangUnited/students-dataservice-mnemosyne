// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: user/user.proto

package user

import (
	common "github.com/NEKETSKY/mnemosyne/pkg/api/common"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request options about composition of user's info
type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WithContacts bool   `protobuf:"varint,1,opt,name=withContacts,proto3" json:"withContacts,omitempty"`
	WithResume   bool   `protobuf:"varint,2,opt,name=withResume,proto3" json:"withResume,omitempty"`
	WithDeleted  bool   `protobuf:"varint,3,opt,name=withDeleted,proto3" json:"withDeleted,omitempty"`
	TeamId       uint32 `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	GroupId      uint32 `protobuf:"varint,5,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *Option) GetWithContacts() bool {
	if x != nil {
		return x.WithContacts
	}
	return false
}

func (x *Option) GetWithResume() bool {
	if x != nil {
		return x.WithResume
	}
	return false
}

func (x *Option) GetWithDeleted() bool {
	if x != nil {
		return x.WithDeleted
	}
	return false
}

func (x *Option) GetTeamId() uint32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *Option) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

// filter by role in da system
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

// filter for info related to user(primary user info, contacts and resume)
type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldName  string `protobuf:"bytes,1,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	FieldValue string `protobuf:"bytes,2,opt,name=fieldValue,proto3" json:"fieldValue,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *Filter) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *Filter) GetFieldValue() string {
	if x != nil {
		return x.FieldValue
	}
	return ""
}

// Summary request options
type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Option *Option `protobuf:"bytes,1,opt,name=option,proto3" json:"option,omitempty"`
	Role   *Role   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Filter *Filter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserRequest) GetOption() *Option {
	if x != nil {
		return x.Option
	}
	return nil
}

func (x *UserRequest) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *UserRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

// message containing info about user's contacts
type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Telegram             string `protobuf:"bytes,2,opt,name=telegram,proto3" json:"telegram,omitempty"`
	Discord              string `protobuf:"bytes,3,opt,name=discord,proto3" json:"discord,omitempty"`
	CommunicationChannel string `protobuf:"bytes,4,opt,name=communicationChannel,proto3" json:"communicationChannel,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *Contact) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Contact) GetTelegram() string {
	if x != nil {
		return x.Telegram
	}
	return ""
}

func (x *Contact) GetDiscord() string {
	if x != nil {
		return x.Discord
	}
	return ""
}

func (x *Contact) GetCommunicationChannel() string {
	if x != nil {
		return x.CommunicationChannel
	}
	return ""
}

// message containing info about user's resume
type Resume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UploadedResume *common.File `protobuf:"bytes,2,opt,name=uploaded_resume,json=uploadedResume,proto3" json:"uploaded_resume,omitempty"`
	MentorsNote    string       `protobuf:"bytes,3,opt,name=mentors_note,json=mentorsNote,proto3" json:"mentors_note,omitempty"`
	Experience     string       `protobuf:"bytes,4,opt,name=experience,proto3" json:"experience,omitempty"`
	Country        string       `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	City           string       `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	TimeZone       string       `protobuf:"bytes,7,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
}

func (x *Resume) Reset() {
	*x = Resume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resume) ProtoMessage() {}

func (x *Resume) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resume.ProtoReflect.Descriptor instead.
func (*Resume) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *Resume) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resume) GetUploadedResume() *common.File {
	if x != nil {
		return x.UploadedResume
	}
	return nil
}

func (x *Resume) GetMentorsNote() string {
	if x != nil {
		return x.MentorsNote
	}
	return ""
}

func (x *Resume) GetExperience() string {
	if x != nil {
		return x.Experience
	}
	return ""
}

func (x *Resume) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Resume) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Resume) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

// The basic message containing the primary information about a user.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LastName     string       `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName    string       `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName   *string      `protobuf:"bytes,4,opt,name=middle_name,json=middleName,proto3,oneof" json:"middle_name,omitempty"`
	Email        string       `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Language     string       `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	EnglishLevel string       `protobuf:"bytes,7,opt,name=english_level,json=englishLevel,proto3" json:"english_level,omitempty"`
	Photo        *common.File `protobuf:"bytes,8,opt,name=photo,proto3" json:"photo,omitempty"`
	Resume       *Resume      `protobuf:"bytes,9,opt,name=resume,proto3" json:"resume,omitempty"`
	Contact      *Contact     `protobuf:"bytes,10,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetMiddleName() string {
	if x != nil && x.MiddleName != nil {
		return *x.MiddleName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *User) GetEnglishLevel() string {
	if x != nil {
		return x.EnglishLevel
	}
	return ""
}

func (x *User) GetPhoto() *common.File {
	if x != nil {
		return x.Photo
	}
	return nil
}

func (x *User) GetResume() *Resume {
	if x != nil {
		return x.Resume
	}
	return nil
}

func (x *User) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

// The basic message containing information about all users from db
type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *Users) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

// Basic request/response containing id of the user
type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{9}
}

func (x *Email) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_user_user_proto protoreflect.FileDescriptor

var file_user_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a,
	0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x77,
	0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x77,
	0x69, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x77, 0x69, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x77,
	0x69, 0x74, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0x1a, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x46, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x3a, 0x3b, 0x92, 0x41, 0x38, 0x0a, 0x36, 0x2a, 0x1b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0xd2, 0x01, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0xd2, 0x01,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0xd2, 0x01, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xcb,
	0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65,
	0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x32, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x3a, 0x46, 0x92, 0x41, 0x43, 0x0a, 0x41, 0x2a, 0x13, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0xd2, 0x01, 0x08, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0xd2, 0x01, 0x07, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x72, 0x64, 0xd2, 0x01, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xc3, 0x02, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0e,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x73, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x3a, 0x64, 0x92, 0x41,
	0x61, 0x0a, 0x5f, 0x2a, 0x12, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xd2, 0x01, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0xd2, 0x01, 0x0c, 0x6d, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0xd2, 0x01, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0xd2, 0x01, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0xd2,
	0x01, 0x04, 0x63, 0x69, 0x74, 0x79, 0xd2, 0x01, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f,
	0x6e, 0x65, 0x22, 0xe2, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x3a, 0x8d, 0x01, 0x92, 0x41, 0x89, 0x01, 0x0a, 0x86,
	0x01, 0x2a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0xd2, 0x01, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2,
	0x01, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0xd2, 0x01, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0xd2,
	0x01, 0x0d, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0xd2,
	0x01, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0xd2, 0x01, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0xd2, 0x01, 0x04, 0x63, 0x69, 0x74, 0x79, 0xd2, 0x01, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x7a, 0x6f, 0x6e, 0x65, 0xd2, 0x01, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0xd2, 0x01, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x20, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x3a, 0x20, 0x92, 0x41, 0x1d, 0x0a, 0x1b, 0x2a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xd2, 0x01, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x1a, 0x92, 0x41, 0x17, 0x0a,
	0x15, 0x2a, 0x0e, 0x49, 0x64, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0xd2, 0x01, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x20, 0x92, 0x41, 0x1d, 0x0a, 0x1b, 0x2a, 0x11, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xd2,
	0x01, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x45, 0x4b, 0x45, 0x54, 0x53, 0x4b, 0x59, 0x2f, 0x6d,
	0x6e, 0x65, 0x6d, 0x6f, 0x73, 0x79, 0x6e, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_user_proto_rawDescOnce sync.Once
	file_user_user_proto_rawDescData = file_user_user_proto_rawDesc
)

func file_user_user_proto_rawDescGZIP() []byte {
	file_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_proto_rawDescData)
	})
	return file_user_user_proto_rawDescData
}

var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_user_proto_goTypes = []interface{}{
	(*Option)(nil),      // 0: user.Option
	(*Role)(nil),        // 1: user.Role
	(*Filter)(nil),      // 2: user.Filter
	(*UserRequest)(nil), // 3: user.UserRequest
	(*Contact)(nil),     // 4: user.Contact
	(*Resume)(nil),      // 5: user.Resume
	(*User)(nil),        // 6: user.User
	(*Users)(nil),       // 7: user.Users
	(*Id)(nil),          // 8: user.Id
	(*Email)(nil),       // 9: user.Email
	(*common.File)(nil), // 10: common.File
}
var file_user_user_proto_depIdxs = []int32{
	0,  // 0: user.UserRequest.option:type_name -> user.Option
	1,  // 1: user.UserRequest.role:type_name -> user.Role
	2,  // 2: user.UserRequest.filter:type_name -> user.Filter
	10, // 3: user.Resume.uploaded_resume:type_name -> common.File
	10, // 4: user.User.photo:type_name -> common.File
	5,  // 5: user.User.resume:type_name -> user.Resume
	4,  // 6: user.User.contact:type_name -> user.Contact
	6,  // 7: user.Users.users:type_name -> user.User
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Option); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resume); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_user_user_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
		MessageInfos:      file_user_user_proto_msgTypes,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_rawDesc = nil
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}
