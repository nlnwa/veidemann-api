// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: contentwriter/v1/contentwriter.proto

package contentwriter

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

// ContentWriterClient is the client API for ContentWriter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentWriterClient interface {
	Write(ctx context.Context, opts ...grpc.CallOption) (ContentWriter_WriteClient, error)
}

type contentWriterClient struct {
	cc grpc.ClientConnInterface
}

func NewContentWriterClient(cc grpc.ClientConnInterface) ContentWriterClient {
	return &contentWriterClient{cc}
}

func (c *contentWriterClient) Write(ctx context.Context, opts ...grpc.CallOption) (ContentWriter_WriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &ContentWriter_ServiceDesc.Streams[0], "/veidemann.api.contentwriter.v1.ContentWriter/write", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentWriterWriteClient{stream}
	return x, nil
}

type ContentWriter_WriteClient interface {
	Send(*WriteRequest) error
	CloseAndRecv() (*WriteReply, error)
	grpc.ClientStream
}

type contentWriterWriteClient struct {
	grpc.ClientStream
}

func (x *contentWriterWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *contentWriterWriteClient) CloseAndRecv() (*WriteReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ContentWriterServer is the server API for ContentWriter service.
// All implementations must embed UnimplementedContentWriterServer
// for forward compatibility
type ContentWriterServer interface {
	Write(ContentWriter_WriteServer) error
	mustEmbedUnimplementedContentWriterServer()
}

// UnimplementedContentWriterServer must be embedded to have forward compatible implementations.
type UnimplementedContentWriterServer struct {
}

func (UnimplementedContentWriterServer) Write(ContentWriter_WriteServer) error {
	return status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedContentWriterServer) mustEmbedUnimplementedContentWriterServer() {}

// UnsafeContentWriterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentWriterServer will
// result in compilation errors.
type UnsafeContentWriterServer interface {
	mustEmbedUnimplementedContentWriterServer()
}

func RegisterContentWriterServer(s grpc.ServiceRegistrar, srv ContentWriterServer) {
	s.RegisterService(&ContentWriter_ServiceDesc, srv)
}

func _ContentWriter_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContentWriterServer).Write(&contentWriterWriteServer{stream})
}

type ContentWriter_WriteServer interface {
	SendAndClose(*WriteReply) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type contentWriterWriteServer struct {
	grpc.ServerStream
}

func (x *contentWriterWriteServer) SendAndClose(m *WriteReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *contentWriterWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ContentWriter_ServiceDesc is the grpc.ServiceDesc for ContentWriter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentWriter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.contentwriter.v1.ContentWriter",
	HandlerType: (*ContentWriterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "write",
			Handler:       _ContentWriter_Write_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "contentwriter/v1/contentwriter.proto",
}
