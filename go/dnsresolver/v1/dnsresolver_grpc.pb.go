// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dnsresolver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DnsResolverClient is the client API for DnsResolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsResolverClient interface {
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveReply, error)
}

type dnsResolverClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsResolverClient(cc grpc.ClientConnInterface) DnsResolverClient {
	return &dnsResolverClient{cc}
}

func (c *dnsResolverClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveReply, error) {
	out := new(ResolveReply)
	err := c.cc.Invoke(ctx, "/veidemann.api.dnsresolver.v1.DnsResolver/resolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsResolverServer is the server API for DnsResolver service.
// All implementations must embed UnimplementedDnsResolverServer
// for forward compatibility
type DnsResolverServer interface {
	Resolve(context.Context, *ResolveRequest) (*ResolveReply, error)
	mustEmbedUnimplementedDnsResolverServer()
}

// UnimplementedDnsResolverServer must be embedded to have forward compatible implementations.
type UnimplementedDnsResolverServer struct {
}

func (UnimplementedDnsResolverServer) Resolve(context.Context, *ResolveRequest) (*ResolveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolve not implemented")
}
func (UnimplementedDnsResolverServer) mustEmbedUnimplementedDnsResolverServer() {}

// UnsafeDnsResolverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DnsResolverServer will
// result in compilation errors.
type UnsafeDnsResolverServer interface {
	mustEmbedUnimplementedDnsResolverServer()
}

func RegisterDnsResolverServer(s grpc.ServiceRegistrar, srv DnsResolverServer) {
	s.RegisterService(&_DnsResolver_serviceDesc, srv)
}

func _DnsResolver_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsResolverServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.dnsresolver.v1.DnsResolver/resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsResolverServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DnsResolver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.dnsresolver.v1.DnsResolver",
	HandlerType: (*DnsResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "resolve",
			Handler:    _DnsResolver_Resolve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dnsresolver/v1/dnsresolver.proto",
}