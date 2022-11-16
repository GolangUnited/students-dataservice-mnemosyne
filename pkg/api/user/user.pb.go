// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: user/user.proto

package user

import (
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

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WithContacts bool `protobuf:"varint,1,opt,name=withContacts,proto3" json:"withContacts,omitempty"`
	WithResume   bool `protobuf:"varint,2,opt,name=withResume,proto3" json:"withResume,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetWithContacts() bool {
	if x != nil {
		return x.WithContacts
	}
	return false
}

func (x *Options) GetWithResume() bool {
	if x != nil {
		return x.WithResume
	}
	return false
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Telegram             string `protobuf:"bytes,1,opt,name=telegram,proto3" json:"telegram,omitempty"`
	Discord              string `protobuf:"bytes,2,opt,name=discord,proto3" json:"discord,omitempty"`
	CommunicationChannel string `protobuf:"bytes,3,opt,name=communicationChannel,proto3" json:"communicationChannel,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{1}
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

// The basic message containing the primary information about a user.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LastName     string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName    string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName   *string  `protobuf:"bytes,4,opt,name=middle_name,json=middleName,proto3,oneof" json:"middle_name,omitempty"`
	Email        string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Language     string   `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	EnglishLevel string   `protobuf:"bytes,7,opt,name=english_level,json=englishLevel,proto3" json:"english_level,omitempty"`
	Photo        string   `protobuf:"bytes,8,opt,name=photo,proto3" json:"photo,omitempty"`
	Experience   string   `protobuf:"bytes,9,opt,name=experience,proto3" json:"experience,omitempty"`
	Country      string   `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`
	City         string   `protobuf:"bytes,11,opt,name=city,proto3" json:"city,omitempty"`
	TimeZone     string   `protobuf:"bytes,12,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	Contact      *Contact `protobuf:"bytes,13,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{2}
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

func (x *User) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *User) GetExperience() string {
	if x != nil {
		return x.Experience
	}
	return ""
}

func (x *User) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *User) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *User) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
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
		mi := &file_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{3}
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
		mi := &file_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateContact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contact *Contact `protobuf:"bytes,1,opt,name=contact,proto3" json:"contact,omitempty"`
	Id      *Id      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateContact) Reset() {
	*x = UpdateContact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateContact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContact) ProtoMessage() {}

func (x *UpdateContact) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateContact.ProtoReflect.Descriptor instead.
func (*UpdateContact) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateContact) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *UpdateContact) GetId() *Id {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_user_user_proto protoreflect.FileDescriptor

var file_user_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x77, 0x69, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x3a, 0x46, 0x92, 0x41,
	0x43, 0x0a, 0x41, 0x2a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xd2, 0x01, 0x08, 0x74, 0x65, 0x6c, 0x65, 0x67,
	0x72, 0x61, 0x6d, 0xd2, 0x01, 0x07, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0xd2, 0x01, 0x14,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x83, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73,
	0x68, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x3a, 0x78, 0x92, 0x41, 0x75, 0x0a, 0x73, 0x2a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xd2, 0x01, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0xd2, 0x01, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0xd2, 0x01, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0xd2, 0x01, 0x0d, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0xd2, 0x01, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0xd2, 0x01, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0xd2, 0x01, 0x04, 0x63, 0x69, 0x74, 0x79, 0xd2, 0x01,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x05, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x20, 0x92, 0x41, 0x1d, 0x0a, 0x1b, 0x2a, 0x11, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xd2,
	0x01, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x1a, 0x92,
	0x41, 0x17, 0x0a, 0x15, 0x2a, 0x0e, 0x49, 0x64, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0xd2, 0x01, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x45, 0x4b, 0x45,
	0x54, 0x53, 0x4b, 0x59, 0x2f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x73, 0x79, 0x6e, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_user_proto_goTypes = []interface{}{
	(*Options)(nil),       // 0: user.Options
	(*Contact)(nil),       // 1: user.Contact
	(*User)(nil),          // 2: user.User
	(*Users)(nil),         // 3: user.Users
	(*Id)(nil),            // 4: user.Id
	(*UpdateContact)(nil), // 5: user.UpdateContact
}
var file_user_user_proto_depIdxs = []int32{
	1, // 0: user.User.contact:type_name -> user.Contact
	2, // 1: user.Users.users:type_name -> user.User
	1, // 2: user.UpdateContact.contact:type_name -> user.Contact
	4, // 3: user.UpdateContact.id:type_name -> user.Id
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
		file_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateContact); i {
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
	file_user_user_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
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
