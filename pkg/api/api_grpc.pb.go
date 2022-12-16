// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api.proto

package api

import (
	context "context"
	certificate "github.com/NEKETSKY/mnemosyne/pkg/api/certificate"
	common "github.com/NEKETSKY/mnemosyne/pkg/api/common"
	helloworld "github.com/NEKETSKY/mnemosyne/pkg/api/helloworld"
	interview "github.com/NEKETSKY/mnemosyne/pkg/api/interview"
	user "github.com/NEKETSKY/mnemosyne/pkg/api/user"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MnemosyneClient is the client API for Mnemosyne service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MnemosyneClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *helloworld.HelloRequest, opts ...grpc.CallOption) (*helloworld.HelloReply, error)
	// Create new user
	CreateUser(ctx context.Context, in *user.User, opts ...grpc.CallOption) (*user.Id, error)
	// Get all existing users
	GetUsers(ctx context.Context, in *user.Options, opts ...grpc.CallOption) (*user.Users, error)
	// Get user by id
	GetUser(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*user.User, error)
	// Update user's data
	UpdateUser(ctx context.Context, in *user.User, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	// Delete user by id
	DeleteUser(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	// Get contact by ID
	GetContact(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*user.Contact, error)
	// Update contact's data
	UpdateContact(ctx context.Context, in *user.Contact, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	// Get resume by ID
	GetResume(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*user.Resume, error)
	// Update resume data
	UpdateResume(ctx context.Context, in *user.Resume, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	//	INTERVIEW
	//
	// Create new interview
	CreateInterview(ctx context.Context, in *interview.Interview, opts ...grpc.CallOption) (*interview.Id, error)
	// Get all existing interviews
	GetInterviews(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*interview.Interview, error)
	// Get interview by id
	GetInterview(ctx context.Context, in *interview.Id, opts ...grpc.CallOption) (*interview.Interview, error)
	// Update interview data
	UpdateInterview(ctx context.Context, in *interview.Interview, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	// Delete contacts by ID
	DeleteContact(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	// Delete resume by ID
	DeleteResume(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	// Delete interview by id
	DeleteInterview(ctx context.Context, in *interview.Id, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	//	Certificate
	//
	// Create new certificate
	CreateCertificate(ctx context.Context, in *certificate.CertificateRequest, opts ...grpc.CallOption) (*certificate.CertificateResponse, error)
	// Get all existing certificates
	GetCertificates(ctx context.Context, in *certificate.Filter, opts ...grpc.CallOption) (*certificate.Certificates, error)
	// Update certificate data
	UpdateCertificate(ctx context.Context, in *certificate.CertificateRequest, opts ...grpc.CallOption) (*common.Empty, error)
	// Deactivate certificate by id
	DeactivateCertificate(ctx context.Context, in *certificate.Id, opts ...grpc.CallOption) (*common.Empty, error)
	// Activate certificate by id
	ActivateCertificate(ctx context.Context, in *certificate.Id, opts ...grpc.CallOption) (*common.Empty, error)
}

type mnemosyneClient struct {
	cc grpc.ClientConnInterface
}

func NewMnemosyneClient(cc grpc.ClientConnInterface) MnemosyneClient {
	return &mnemosyneClient{cc}
}

func (c *mnemosyneClient) SayHello(ctx context.Context, in *helloworld.HelloRequest, opts ...grpc.CallOption) (*helloworld.HelloReply, error) {
	out := new(helloworld.HelloReply)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) CreateUser(ctx context.Context, in *user.User, opts ...grpc.CallOption) (*user.Id, error) {
	out := new(user.Id)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) GetUsers(ctx context.Context, in *user.Options, opts ...grpc.CallOption) (*user.Users, error) {
	out := new(user.Users)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) GetUser(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*user.User, error) {
	out := new(user.User)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) UpdateUser(ctx context.Context, in *user.User, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) DeleteUser(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) GetContact(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*user.Contact, error) {
	out := new(user.Contact)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/GetContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) UpdateContact(ctx context.Context, in *user.Contact, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/UpdateContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) GetResume(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*user.Resume, error) {
	out := new(user.Resume)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/GetResume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) UpdateResume(ctx context.Context, in *user.Resume, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/UpdateResume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) CreateInterview(ctx context.Context, in *interview.Interview, opts ...grpc.CallOption) (*interview.Id, error) {
	out := new(interview.Id)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/CreateInterview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) GetInterviews(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*interview.Interview, error) {
	out := new(interview.Interview)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/GetInterviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) GetInterview(ctx context.Context, in *interview.Id, opts ...grpc.CallOption) (*interview.Interview, error) {
	out := new(interview.Interview)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/GetInterview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) UpdateInterview(ctx context.Context, in *interview.Interview, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/UpdateInterview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) DeleteContact(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/DeleteContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) DeleteResume(ctx context.Context, in *user.Id, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/DeleteResume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) DeleteInterview(ctx context.Context, in *interview.Id, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/DeleteInterview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) CreateCertificate(ctx context.Context, in *certificate.CertificateRequest, opts ...grpc.CallOption) (*certificate.CertificateResponse, error) {
	out := new(certificate.CertificateResponse)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/CreateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) GetCertificates(ctx context.Context, in *certificate.Filter, opts ...grpc.CallOption) (*certificate.Certificates, error) {
	out := new(certificate.Certificates)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/GetCertificates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) UpdateCertificate(ctx context.Context, in *certificate.CertificateRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/UpdateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) DeactivateCertificate(ctx context.Context, in *certificate.Id, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/DeactivateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mnemosyneClient) ActivateCertificate(ctx context.Context, in *certificate.Id, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/api.Mnemosyne/ActivateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MnemosyneServer is the server API for Mnemosyne service.
// All implementations should embed UnimplementedMnemosyneServer
// for forward compatibility
type MnemosyneServer interface {
	// Sends a greeting
	SayHello(context.Context, *helloworld.HelloRequest) (*helloworld.HelloReply, error)
	// Create new user
	CreateUser(context.Context, *user.User) (*user.Id, error)
	// Get all existing users
	GetUsers(context.Context, *user.Options) (*user.Users, error)
	// Get user by id
	GetUser(context.Context, *user.Id) (*user.User, error)
	// Update user's data
	UpdateUser(context.Context, *user.User) (*wrapperspb.BoolValue, error)
	// Delete user by id
	DeleteUser(context.Context, *user.Id) (*wrapperspb.BoolValue, error)
	// Get contact by ID
	GetContact(context.Context, *user.Id) (*user.Contact, error)
	// Update contact's data
	UpdateContact(context.Context, *user.Contact) (*wrapperspb.BoolValue, error)
	// Get resume by ID
	GetResume(context.Context, *user.Id) (*user.Resume, error)
	// Update resume data
	UpdateResume(context.Context, *user.Resume) (*wrapperspb.BoolValue, error)
	//	INTERVIEW
	//
	// Create new interview
	CreateInterview(context.Context, *interview.Interview) (*interview.Id, error)
	// Get all existing interviews
	GetInterviews(context.Context, *common.Empty) (*interview.Interview, error)
	// Get interview by id
	GetInterview(context.Context, *interview.Id) (*interview.Interview, error)
	// Update interview data
	UpdateInterview(context.Context, *interview.Interview) (*wrapperspb.BoolValue, error)
	// Delete contacts by ID
	DeleteContact(context.Context, *user.Id) (*wrapperspb.BoolValue, error)
	// Delete resume by ID
	DeleteResume(context.Context, *user.Id) (*wrapperspb.BoolValue, error)
	// Delete interview by id
	DeleteInterview(context.Context, *interview.Id) (*wrapperspb.BoolValue, error)
	//	Certificate
	//
	// Create new certificate
	CreateCertificate(context.Context, *certificate.CertificateRequest) (*certificate.CertificateResponse, error)
	// Get all existing certificates
	GetCertificates(context.Context, *certificate.Filter) (*certificate.Certificates, error)
	// Update certificate data
	UpdateCertificate(context.Context, *certificate.CertificateRequest) (*common.Empty, error)
	// Deactivate certificate by id
	DeactivateCertificate(context.Context, *certificate.Id) (*common.Empty, error)
	// Activate certificate by id
	ActivateCertificate(context.Context, *certificate.Id) (*common.Empty, error)
}

// UnimplementedMnemosyneServer should be embedded to have forward compatible implementations.
type UnimplementedMnemosyneServer struct {
}

func (UnimplementedMnemosyneServer) SayHello(context.Context, *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedMnemosyneServer) CreateUser(context.Context, *user.User) (*user.Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMnemosyneServer) GetUsers(context.Context, *user.Options) (*user.Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedMnemosyneServer) GetUser(context.Context, *user.Id) (*user.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedMnemosyneServer) UpdateUser(context.Context, *user.User) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedMnemosyneServer) DeleteUser(context.Context, *user.Id) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedMnemosyneServer) GetContact(context.Context, *user.Id) (*user.Contact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContact not implemented")
}
func (UnimplementedMnemosyneServer) UpdateContact(context.Context, *user.Contact) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContact not implemented")
}
func (UnimplementedMnemosyneServer) GetResume(context.Context, *user.Id) (*user.Resume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResume not implemented")
}
func (UnimplementedMnemosyneServer) UpdateResume(context.Context, *user.Resume) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResume not implemented")
}
func (UnimplementedMnemosyneServer) CreateInterview(context.Context, *interview.Interview) (*interview.Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInterview not implemented")
}
func (UnimplementedMnemosyneServer) GetInterviews(context.Context, *common.Empty) (*interview.Interview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInterviews not implemented")
}
func (UnimplementedMnemosyneServer) GetInterview(context.Context, *interview.Id) (*interview.Interview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInterview not implemented")
}
func (UnimplementedMnemosyneServer) UpdateInterview(context.Context, *interview.Interview) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInterview not implemented")
}
func (UnimplementedMnemosyneServer) DeleteContact(context.Context, *user.Id) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContact not implemented")
}
func (UnimplementedMnemosyneServer) DeleteResume(context.Context, *user.Id) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResume not implemented")
}
func (UnimplementedMnemosyneServer) DeleteInterview(context.Context, *interview.Id) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInterview not implemented")
}
func (UnimplementedMnemosyneServer) CreateCertificate(context.Context, *certificate.CertificateRequest) (*certificate.CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}
func (UnimplementedMnemosyneServer) GetCertificates(context.Context, *certificate.Filter) (*certificate.Certificates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificates not implemented")
}
func (UnimplementedMnemosyneServer) UpdateCertificate(context.Context, *certificate.CertificateRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCertificate not implemented")
}
func (UnimplementedMnemosyneServer) DeactivateCertificate(context.Context, *certificate.Id) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateCertificate not implemented")
}
func (UnimplementedMnemosyneServer) ActivateCertificate(context.Context, *certificate.Id) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateCertificate not implemented")
}

// UnsafeMnemosyneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MnemosyneServer will
// result in compilation errors.
type UnsafeMnemosyneServer interface {
	mustEmbedUnimplementedMnemosyneServer()
}

func RegisterMnemosyneServer(s grpc.ServiceRegistrar, srv MnemosyneServer) {
	s.RegisterService(&Mnemosyne_ServiceDesc, srv)
}

func _Mnemosyne_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helloworld.HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).SayHello(ctx, req.(*helloworld.HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).CreateUser(ctx, req.(*user.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).GetUsers(ctx, req.(*user.Options))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).GetUser(ctx, req.(*user.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).UpdateUser(ctx, req.(*user.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).DeleteUser(ctx, req.(*user.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_GetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).GetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/GetContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).GetContact(ctx, req.(*user.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_UpdateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).UpdateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/UpdateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).UpdateContact(ctx, req.(*user.Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_GetResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).GetResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/GetResume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).GetResume(ctx, req.(*user.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_UpdateResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.Resume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).UpdateResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/UpdateResume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).UpdateResume(ctx, req.(*user.Resume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_CreateInterview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(interview.Interview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).CreateInterview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/CreateInterview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).CreateInterview(ctx, req.(*interview.Interview))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_GetInterviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).GetInterviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/GetInterviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).GetInterviews(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_GetInterview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(interview.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).GetInterview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/GetInterview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).GetInterview(ctx, req.(*interview.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_UpdateInterview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(interview.Interview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).UpdateInterview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/UpdateInterview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).UpdateInterview(ctx, req.(*interview.Interview))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_DeleteContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).DeleteContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/DeleteContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).DeleteContact(ctx, req.(*user.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_DeleteResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).DeleteResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/DeleteResume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).DeleteResume(ctx, req.(*user.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_DeleteInterview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(interview.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).DeleteInterview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/DeleteInterview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).DeleteInterview(ctx, req.(*interview.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(certificate.CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/CreateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).CreateCertificate(ctx, req.(*certificate.CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_GetCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(certificate.Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).GetCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/GetCertificates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).GetCertificates(ctx, req.(*certificate.Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_UpdateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(certificate.CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).UpdateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/UpdateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).UpdateCertificate(ctx, req.(*certificate.CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_DeactivateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(certificate.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).DeactivateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/DeactivateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).DeactivateCertificate(ctx, req.(*certificate.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mnemosyne_ActivateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(certificate.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MnemosyneServer).ActivateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Mnemosyne/ActivateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MnemosyneServer).ActivateCertificate(ctx, req.(*certificate.Id))
	}
	return interceptor(ctx, in, info, handler)
}

// Mnemosyne_ServiceDesc is the grpc.ServiceDesc for Mnemosyne service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mnemosyne_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Mnemosyne",
	HandlerType: (*MnemosyneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Mnemosyne_SayHello_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Mnemosyne_CreateUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Mnemosyne_GetUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Mnemosyne_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Mnemosyne_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Mnemosyne_DeleteUser_Handler,
		},
		{
			MethodName: "GetContact",
			Handler:    _Mnemosyne_GetContact_Handler,
		},
		{
			MethodName: "UpdateContact",
			Handler:    _Mnemosyne_UpdateContact_Handler,
		},
		{
			MethodName: "GetResume",
			Handler:    _Mnemosyne_GetResume_Handler,
		},
		{
			MethodName: "UpdateResume",
			Handler:    _Mnemosyne_UpdateResume_Handler,
		},
		{
			MethodName: "CreateInterview",
			Handler:    _Mnemosyne_CreateInterview_Handler,
		},
		{
			MethodName: "GetInterviews",
			Handler:    _Mnemosyne_GetInterviews_Handler,
		},
		{
			MethodName: "GetInterview",
			Handler:    _Mnemosyne_GetInterview_Handler,
		},
		{
			MethodName: "UpdateInterview",
			Handler:    _Mnemosyne_UpdateInterview_Handler,
		},
		{
			MethodName: "DeleteContact",
			Handler:    _Mnemosyne_DeleteContact_Handler,
		},
		{
			MethodName: "DeleteResume",
			Handler:    _Mnemosyne_DeleteResume_Handler,
		},
		{
			MethodName: "DeleteInterview",
			Handler:    _Mnemosyne_DeleteInterview_Handler,
		},
		{
			MethodName: "CreateCertificate",
			Handler:    _Mnemosyne_CreateCertificate_Handler,
		},
		{
			MethodName: "GetCertificates",
			Handler:    _Mnemosyne_GetCertificates_Handler,
		},
		{
			MethodName: "UpdateCertificate",
			Handler:    _Mnemosyne_UpdateCertificate_Handler,
		},
		{
			MethodName: "DeactivateCertificate",
			Handler:    _Mnemosyne_DeactivateCertificate_Handler,
		},
		{
			MethodName: "ActivateCertificate",
			Handler:    _Mnemosyne_ActivateCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
