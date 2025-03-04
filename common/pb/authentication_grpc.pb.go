// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: authentication.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthenticationService_SignInWithEmail_FullMethodName  = "/pb.AuthenticationService/SignInWithEmail"
	AuthenticationService_SignInWithPhone_FullMethodName  = "/pb.AuthenticationService/SignInWithPhone"
	AuthenticationService_SignInWithGoogle_FullMethodName = "/pb.AuthenticationService/SignInWithGoogle"
	AuthenticationService_GetAuthById_FullMethodName      = "/pb.AuthenticationService/GetAuthById"
	AuthenticationService_GetAuth_FullMethodName          = "/pb.AuthenticationService/GetAuth"
	AuthenticationService_RefreshToken_FullMethodName     = "/pb.AuthenticationService/RefreshToken"
	AuthenticationService_LogOut_FullMethodName           = "/pb.AuthenticationService/LogOut"
)

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	SignInWithEmail(ctx context.Context, in *SignInWithEmailReq, opts ...grpc.CallOption) (*SignInRes, error)
	SignInWithPhone(ctx context.Context, in *SignInWithPhoneReq, opts ...grpc.CallOption) (*SignInRes, error)
	SignInWithGoogle(ctx context.Context, in *SignInWithGoogleReq, opts ...grpc.CallOption) (*SignInRes, error)
	GetAuthById(ctx context.Context, in *GetAuthByIdReq, opts ...grpc.CallOption) (*Auth, error)
	GetAuth(ctx context.Context, in *GetAuthReq, opts ...grpc.CallOption) (*Auth, error)
	// TODO: these rpc's should be protected
	RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*SignInRes, error)
	LogOut(ctx context.Context, in *LogOutReq, opts ...grpc.CallOption) (*SignInRes, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) SignInWithEmail(ctx context.Context, in *SignInWithEmailReq, opts ...grpc.CallOption) (*SignInRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInRes)
	err := c.cc.Invoke(ctx, AuthenticationService_SignInWithEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) SignInWithPhone(ctx context.Context, in *SignInWithPhoneReq, opts ...grpc.CallOption) (*SignInRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInRes)
	err := c.cc.Invoke(ctx, AuthenticationService_SignInWithPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) SignInWithGoogle(ctx context.Context, in *SignInWithGoogleReq, opts ...grpc.CallOption) (*SignInRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInRes)
	err := c.cc.Invoke(ctx, AuthenticationService_SignInWithGoogle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) GetAuthById(ctx context.Context, in *GetAuthByIdReq, opts ...grpc.CallOption) (*Auth, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Auth)
	err := c.cc.Invoke(ctx, AuthenticationService_GetAuthById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) GetAuth(ctx context.Context, in *GetAuthReq, opts ...grpc.CallOption) (*Auth, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Auth)
	err := c.cc.Invoke(ctx, AuthenticationService_GetAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*SignInRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInRes)
	err := c.cc.Invoke(ctx, AuthenticationService_RefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) LogOut(ctx context.Context, in *LogOutReq, opts ...grpc.CallOption) (*SignInRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInRes)
	err := c.cc.Invoke(ctx, AuthenticationService_LogOut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
// All implementations must embed UnimplementedAuthenticationServiceServer
// for forward compatibility.
type AuthenticationServiceServer interface {
	SignInWithEmail(context.Context, *SignInWithEmailReq) (*SignInRes, error)
	SignInWithPhone(context.Context, *SignInWithPhoneReq) (*SignInRes, error)
	SignInWithGoogle(context.Context, *SignInWithGoogleReq) (*SignInRes, error)
	GetAuthById(context.Context, *GetAuthByIdReq) (*Auth, error)
	GetAuth(context.Context, *GetAuthReq) (*Auth, error)
	// TODO: these rpc's should be protected
	RefreshToken(context.Context, *RefreshTokenReq) (*SignInRes, error)
	LogOut(context.Context, *LogOutReq) (*SignInRes, error)
	mustEmbedUnimplementedAuthenticationServiceServer()
}

// UnimplementedAuthenticationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthenticationServiceServer struct{}

func (UnimplementedAuthenticationServiceServer) SignInWithEmail(context.Context, *SignInWithEmailReq) (*SignInRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInWithEmail not implemented")
}
func (UnimplementedAuthenticationServiceServer) SignInWithPhone(context.Context, *SignInWithPhoneReq) (*SignInRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInWithPhone not implemented")
}
func (UnimplementedAuthenticationServiceServer) SignInWithGoogle(context.Context, *SignInWithGoogleReq) (*SignInRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInWithGoogle not implemented")
}
func (UnimplementedAuthenticationServiceServer) GetAuthById(context.Context, *GetAuthByIdReq) (*Auth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthById not implemented")
}
func (UnimplementedAuthenticationServiceServer) GetAuth(context.Context, *GetAuthReq) (*Auth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuth not implemented")
}
func (UnimplementedAuthenticationServiceServer) RefreshToken(context.Context, *RefreshTokenReq) (*SignInRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthenticationServiceServer) LogOut(context.Context, *LogOutReq) (*SignInRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedAuthenticationServiceServer) mustEmbedUnimplementedAuthenticationServiceServer() {}
func (UnimplementedAuthenticationServiceServer) testEmbeddedByValue()                               {}

// UnsafeAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServiceServer will
// result in compilation errors.
type UnsafeAuthenticationServiceServer interface {
	mustEmbedUnimplementedAuthenticationServiceServer()
}

func RegisterAuthenticationServiceServer(s grpc.ServiceRegistrar, srv AuthenticationServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthenticationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthenticationService_ServiceDesc, srv)
}

func _AuthenticationService_SignInWithEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInWithEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).SignInWithEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_SignInWithEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).SignInWithEmail(ctx, req.(*SignInWithEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_SignInWithPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInWithPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).SignInWithPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_SignInWithPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).SignInWithPhone(ctx, req.(*SignInWithPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_SignInWithGoogle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInWithGoogleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).SignInWithGoogle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_SignInWithGoogle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).SignInWithGoogle(ctx, req.(*SignInWithGoogleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_GetAuthById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GetAuthById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_GetAuthById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GetAuthById(ctx, req.(*GetAuthByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_GetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_GetAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GetAuth(ctx, req.(*GetAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).RefreshToken(ctx, req.(*RefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogOutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_LogOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).LogOut(ctx, req.(*LogOutReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationService_ServiceDesc is the grpc.ServiceDesc for AuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignInWithEmail",
			Handler:    _AuthenticationService_SignInWithEmail_Handler,
		},
		{
			MethodName: "SignInWithPhone",
			Handler:    _AuthenticationService_SignInWithPhone_Handler,
		},
		{
			MethodName: "SignInWithGoogle",
			Handler:    _AuthenticationService_SignInWithGoogle_Handler,
		},
		{
			MethodName: "GetAuthById",
			Handler:    _AuthenticationService_GetAuthById_Handler,
		},
		{
			MethodName: "GetAuth",
			Handler:    _AuthenticationService_GetAuth_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthenticationService_RefreshToken_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _AuthenticationService_LogOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}
