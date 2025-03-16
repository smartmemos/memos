// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/v1/user.proto

package v1

import (
	context "context"
	user "github.com/smartmemos/memos/internal/proto/model/user"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_CreateUser_FullMethodName        = "/api.v1.UserService/CreateUser"
	UserService_GetUserSetting_FullMethodName    = "/api.v1.UserService/GetUserSetting"
	UserService_UpdateUserSetting_FullMethodName = "/api.v1.UserService/UpdateUserSetting"
	UserService_ListAllUserStats_FullMethodName  = "/api.v1.UserService/ListAllUserStats"
	UserService_GetUserStats_FullMethodName      = "/api.v1.UserService/GetUserStats"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*user.User, error)
	// GetUserSetting gets the setting of a user.
	GetUserSetting(ctx context.Context, in *GetUserSettingRequest, opts ...grpc.CallOption) (*user.Setting, error)
	// UpdateUserSetting updates the setting of a user.
	UpdateUserSetting(ctx context.Context, in *UpdateUserSettingRequest, opts ...grpc.CallOption) (*user.Setting, error)
	// ListAllUserStats returns all user stats.
	ListAllUserStats(ctx context.Context, in *ListAllUserStatsRequest, opts ...grpc.CallOption) (*ListAllUserStatsResponse, error)
	// GetUserStats returns the stats of a user.
	GetUserStats(ctx context.Context, in *GetUserStatsRequest, opts ...grpc.CallOption) (*user.Stats, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*user.User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(user.User)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserSetting(ctx context.Context, in *GetUserSettingRequest, opts ...grpc.CallOption) (*user.Setting, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(user.Setting)
	err := c.cc.Invoke(ctx, UserService_GetUserSetting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserSetting(ctx context.Context, in *UpdateUserSettingRequest, opts ...grpc.CallOption) (*user.Setting, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(user.Setting)
	err := c.cc.Invoke(ctx, UserService_UpdateUserSetting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListAllUserStats(ctx context.Context, in *ListAllUserStatsRequest, opts ...grpc.CallOption) (*ListAllUserStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAllUserStatsResponse)
	err := c.cc.Invoke(ctx, UserService_ListAllUserStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserStats(ctx context.Context, in *GetUserStatsRequest, opts ...grpc.CallOption) (*user.Stats, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(user.Stats)
	err := c.cc.Invoke(ctx, UserService_GetUserStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	// CreateUser creates a new user.
	CreateUser(context.Context, *CreateUserRequest) (*user.User, error)
	// GetUserSetting gets the setting of a user.
	GetUserSetting(context.Context, *GetUserSettingRequest) (*user.Setting, error)
	// UpdateUserSetting updates the setting of a user.
	UpdateUserSetting(context.Context, *UpdateUserSettingRequest) (*user.Setting, error)
	// ListAllUserStats returns all user stats.
	ListAllUserStats(context.Context, *ListAllUserStatsRequest) (*ListAllUserStatsResponse, error)
	// GetUserStats returns the stats of a user.
	GetUserStats(context.Context, *GetUserStatsRequest) (*user.Stats, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*user.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserSetting(context.Context, *GetUserSettingRequest) (*user.Setting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSetting not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserSetting(context.Context, *UpdateUserSettingRequest) (*user.Setting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserSetting not implemented")
}
func (UnimplementedUserServiceServer) ListAllUserStats(context.Context, *ListAllUserStatsRequest) (*ListAllUserStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllUserStats not implemented")
}
func (UnimplementedUserServiceServer) GetUserStats(context.Context, *GetUserStatsRequest) (*user.Stats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStats not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserSetting(ctx, req.(*GetUserSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserSetting(ctx, req.(*UpdateUserSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListAllUserStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllUserStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListAllUserStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListAllUserStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListAllUserStats(ctx, req.(*ListAllUserStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserStats(ctx, req.(*GetUserStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserSetting",
			Handler:    _UserService_GetUserSetting_Handler,
		},
		{
			MethodName: "UpdateUserSetting",
			Handler:    _UserService_UpdateUserSetting_Handler,
		},
		{
			MethodName: "ListAllUserStats",
			Handler:    _UserService_ListAllUserStats_Handler,
		},
		{
			MethodName: "GetUserStats",
			Handler:    _UserService_GetUserStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/user.proto",
}
