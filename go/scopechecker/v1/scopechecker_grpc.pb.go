// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: scopechecker/v1/scopechecker.proto

package scopechecker

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ScopesCheckerServiceClient is the client API for ScopesCheckerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScopesCheckerServiceClient interface {
	// Check if URI is in scope for this crawl
	ScopeCheck(ctx context.Context, in *ScopeCheckRequest, opts ...grpc.CallOption) (*ScopeCheckResponse, error)
}

type scopesCheckerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScopesCheckerServiceClient(cc grpc.ClientConnInterface) ScopesCheckerServiceClient {
	return &scopesCheckerServiceClient{cc}
}

func (c *scopesCheckerServiceClient) ScopeCheck(ctx context.Context, in *ScopeCheckRequest, opts ...grpc.CallOption) (*ScopeCheckResponse, error) {
	out := new(ScopeCheckResponse)
	err := c.cc.Invoke(ctx, "/veidemann.api.scopechecker.v1.ScopesCheckerService/ScopeCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScopesCheckerServiceServer is the server API for ScopesCheckerService service.
// All implementations must embed UnimplementedScopesCheckerServiceServer
// for forward compatibility
type ScopesCheckerServiceServer interface {
	// Check if URI is in scope for this crawl
	ScopeCheck(context.Context, *ScopeCheckRequest) (*ScopeCheckResponse, error)
	mustEmbedUnimplementedScopesCheckerServiceServer()
}

// UnimplementedScopesCheckerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScopesCheckerServiceServer struct {
}

func (UnimplementedScopesCheckerServiceServer) ScopeCheck(context.Context, *ScopeCheckRequest) (*ScopeCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScopeCheck not implemented")
}
func (UnimplementedScopesCheckerServiceServer) mustEmbedUnimplementedScopesCheckerServiceServer() {}

// UnsafeScopesCheckerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScopesCheckerServiceServer will
// result in compilation errors.
type UnsafeScopesCheckerServiceServer interface {
	mustEmbedUnimplementedScopesCheckerServiceServer()
}

func RegisterScopesCheckerServiceServer(s grpc.ServiceRegistrar, srv ScopesCheckerServiceServer) {
	s.RegisterService(&ScopesCheckerService_ServiceDesc, srv)
}

func _ScopesCheckerService_ScopeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScopeCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopesCheckerServiceServer).ScopeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.scopechecker.v1.ScopesCheckerService/ScopeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopesCheckerServiceServer).ScopeCheck(ctx, req.(*ScopeCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScopesCheckerService_ServiceDesc is the grpc.ServiceDesc for ScopesCheckerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScopesCheckerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.scopechecker.v1.ScopesCheckerService",
	HandlerType: (*ScopesCheckerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScopeCheck",
			Handler:    _ScopesCheckerService_ScopeCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scopechecker/v1/scopechecker.proto",
}
