// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: robotsevaluator/v2/robotsevaluator.proto

package robotsevaluator

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

// RobotsEvaluatorClient is the client API for RobotsEvaluator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotsEvaluatorClient interface {
	IsAllowed(ctx context.Context, in *IsAllowedRequest, opts ...grpc.CallOption) (*IsAllowedReply, error)
}

type robotsEvaluatorClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotsEvaluatorClient(cc grpc.ClientConnInterface) RobotsEvaluatorClient {
	return &robotsEvaluatorClient{cc}
}

func (c *robotsEvaluatorClient) IsAllowed(ctx context.Context, in *IsAllowedRequest, opts ...grpc.CallOption) (*IsAllowedReply, error) {
	out := new(IsAllowedReply)
	err := c.cc.Invoke(ctx, "/veidemann.api.robotsevaluator.v2.RobotsEvaluator/isAllowed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotsEvaluatorServer is the server API for RobotsEvaluator service.
// All implementations must embed UnimplementedRobotsEvaluatorServer
// for forward compatibility
type RobotsEvaluatorServer interface {
	IsAllowed(context.Context, *IsAllowedRequest) (*IsAllowedReply, error)
	mustEmbedUnimplementedRobotsEvaluatorServer()
}

// UnimplementedRobotsEvaluatorServer must be embedded to have forward compatible implementations.
type UnimplementedRobotsEvaluatorServer struct {
}

func (UnimplementedRobotsEvaluatorServer) IsAllowed(context.Context, *IsAllowedRequest) (*IsAllowedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAllowed not implemented")
}
func (UnimplementedRobotsEvaluatorServer) mustEmbedUnimplementedRobotsEvaluatorServer() {}

// UnsafeRobotsEvaluatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotsEvaluatorServer will
// result in compilation errors.
type UnsafeRobotsEvaluatorServer interface {
	mustEmbedUnimplementedRobotsEvaluatorServer()
}

func RegisterRobotsEvaluatorServer(s grpc.ServiceRegistrar, srv RobotsEvaluatorServer) {
	s.RegisterService(&RobotsEvaluator_ServiceDesc, srv)
}

func _RobotsEvaluator_IsAllowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAllowedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsEvaluatorServer).IsAllowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.robotsevaluator.v2.RobotsEvaluator/isAllowed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsEvaluatorServer).IsAllowed(ctx, req.(*IsAllowedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RobotsEvaluator_ServiceDesc is the grpc.ServiceDesc for RobotsEvaluator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RobotsEvaluator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.robotsevaluator.v2.RobotsEvaluator",
	HandlerType: (*RobotsEvaluatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "isAllowed",
			Handler:    _RobotsEvaluator_IsAllowed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "robotsevaluator/v2/robotsevaluator.proto",
}
