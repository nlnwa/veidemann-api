// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: config/v1/config.proto

package config

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConfigClient is the client API for Config service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigClient interface {
	// Get a config object by ID
	GetConfigObject(ctx context.Context, in *ConfigRef, opts ...grpc.CallOption) (*ConfigObject, error)
	// List a set of config objects
	ListConfigObjects(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Config_ListConfigObjectsClient, error)
	// Count config objects
	CountConfigObjects(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCountResponse, error)
	// Save a config object
	SaveConfigObject(ctx context.Context, in *ConfigObject, opts ...grpc.CallOption) (*ConfigObject, error)
	// Update config objects
	UpdateConfigObjects(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete a config object
	DeleteConfigObject(ctx context.Context, in *ConfigObject, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetLabelKeys(ctx context.Context, in *GetLabelKeysRequest, opts ...grpc.CallOption) (*LabelKeysResponse, error)
	GetLogConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LogLevels, error)
	SaveLogConfig(ctx context.Context, in *LogLevels, opts ...grpc.CallOption) (*LogLevels, error)
	GetScriptAnnotations(ctx context.Context, in *GetScriptAnnotationsRequest, opts ...grpc.CallOption) (*GetScriptAnnotationsResponse, error)
}

type configClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigClient(cc grpc.ClientConnInterface) ConfigClient {
	return &configClient{cc}
}

func (c *configClient) GetConfigObject(ctx context.Context, in *ConfigRef, opts ...grpc.CallOption) (*ConfigObject, error) {
	out := new(ConfigObject)
	err := c.cc.Invoke(ctx, "/veidemann.api.config.v1.Config/GetConfigObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) ListConfigObjects(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Config_ListConfigObjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Config_ServiceDesc.Streams[0], "/veidemann.api.config.v1.Config/ListConfigObjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &configListConfigObjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Config_ListConfigObjectsClient interface {
	Recv() (*ConfigObject, error)
	grpc.ClientStream
}

type configListConfigObjectsClient struct {
	grpc.ClientStream
}

func (x *configListConfigObjectsClient) Recv() (*ConfigObject, error) {
	m := new(ConfigObject)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *configClient) CountConfigObjects(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCountResponse, error) {
	out := new(ListCountResponse)
	err := c.cc.Invoke(ctx, "/veidemann.api.config.v1.Config/CountConfigObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) SaveConfigObject(ctx context.Context, in *ConfigObject, opts ...grpc.CallOption) (*ConfigObject, error) {
	out := new(ConfigObject)
	err := c.cc.Invoke(ctx, "/veidemann.api.config.v1.Config/SaveConfigObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) UpdateConfigObjects(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/veidemann.api.config.v1.Config/UpdateConfigObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) DeleteConfigObject(ctx context.Context, in *ConfigObject, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/veidemann.api.config.v1.Config/DeleteConfigObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) GetLabelKeys(ctx context.Context, in *GetLabelKeysRequest, opts ...grpc.CallOption) (*LabelKeysResponse, error) {
	out := new(LabelKeysResponse)
	err := c.cc.Invoke(ctx, "/veidemann.api.config.v1.Config/GetLabelKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) GetLogConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LogLevels, error) {
	out := new(LogLevels)
	err := c.cc.Invoke(ctx, "/veidemann.api.config.v1.Config/GetLogConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) SaveLogConfig(ctx context.Context, in *LogLevels, opts ...grpc.CallOption) (*LogLevels, error) {
	out := new(LogLevels)
	err := c.cc.Invoke(ctx, "/veidemann.api.config.v1.Config/SaveLogConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) GetScriptAnnotations(ctx context.Context, in *GetScriptAnnotationsRequest, opts ...grpc.CallOption) (*GetScriptAnnotationsResponse, error) {
	out := new(GetScriptAnnotationsResponse)
	err := c.cc.Invoke(ctx, "/veidemann.api.config.v1.Config/GetScriptAnnotations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServer is the server API for Config service.
// All implementations must embed UnimplementedConfigServer
// for forward compatibility
type ConfigServer interface {
	// Get a config object by ID
	GetConfigObject(context.Context, *ConfigRef) (*ConfigObject, error)
	// List a set of config objects
	ListConfigObjects(*ListRequest, Config_ListConfigObjectsServer) error
	// Count config objects
	CountConfigObjects(context.Context, *ListRequest) (*ListCountResponse, error)
	// Save a config object
	SaveConfigObject(context.Context, *ConfigObject) (*ConfigObject, error)
	// Update config objects
	UpdateConfigObjects(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete a config object
	DeleteConfigObject(context.Context, *ConfigObject) (*DeleteResponse, error)
	GetLabelKeys(context.Context, *GetLabelKeysRequest) (*LabelKeysResponse, error)
	GetLogConfig(context.Context, *emptypb.Empty) (*LogLevels, error)
	SaveLogConfig(context.Context, *LogLevels) (*LogLevels, error)
	GetScriptAnnotations(context.Context, *GetScriptAnnotationsRequest) (*GetScriptAnnotationsResponse, error)
	mustEmbedUnimplementedConfigServer()
}

// UnimplementedConfigServer must be embedded to have forward compatible implementations.
type UnimplementedConfigServer struct {
}

func (UnimplementedConfigServer) GetConfigObject(context.Context, *ConfigRef) (*ConfigObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigObject not implemented")
}
func (UnimplementedConfigServer) ListConfigObjects(*ListRequest, Config_ListConfigObjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListConfigObjects not implemented")
}
func (UnimplementedConfigServer) CountConfigObjects(context.Context, *ListRequest) (*ListCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountConfigObjects not implemented")
}
func (UnimplementedConfigServer) SaveConfigObject(context.Context, *ConfigObject) (*ConfigObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveConfigObject not implemented")
}
func (UnimplementedConfigServer) UpdateConfigObjects(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfigObjects not implemented")
}
func (UnimplementedConfigServer) DeleteConfigObject(context.Context, *ConfigObject) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfigObject not implemented")
}
func (UnimplementedConfigServer) GetLabelKeys(context.Context, *GetLabelKeysRequest) (*LabelKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabelKeys not implemented")
}
func (UnimplementedConfigServer) GetLogConfig(context.Context, *emptypb.Empty) (*LogLevels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogConfig not implemented")
}
func (UnimplementedConfigServer) SaveLogConfig(context.Context, *LogLevels) (*LogLevels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLogConfig not implemented")
}
func (UnimplementedConfigServer) GetScriptAnnotations(context.Context, *GetScriptAnnotationsRequest) (*GetScriptAnnotationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScriptAnnotations not implemented")
}
func (UnimplementedConfigServer) mustEmbedUnimplementedConfigServer() {}

// UnsafeConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServer will
// result in compilation errors.
type UnsafeConfigServer interface {
	mustEmbedUnimplementedConfigServer()
}

func RegisterConfigServer(s grpc.ServiceRegistrar, srv ConfigServer) {
	s.RegisterService(&Config_ServiceDesc, srv)
}

func _Config_GetConfigObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).GetConfigObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.config.v1.Config/GetConfigObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).GetConfigObject(ctx, req.(*ConfigRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_ListConfigObjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigServer).ListConfigObjects(m, &configListConfigObjectsServer{stream})
}

type Config_ListConfigObjectsServer interface {
	Send(*ConfigObject) error
	grpc.ServerStream
}

type configListConfigObjectsServer struct {
	grpc.ServerStream
}

func (x *configListConfigObjectsServer) Send(m *ConfigObject) error {
	return x.ServerStream.SendMsg(m)
}

func _Config_CountConfigObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).CountConfigObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.config.v1.Config/CountConfigObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).CountConfigObjects(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_SaveConfigObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).SaveConfigObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.config.v1.Config/SaveConfigObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).SaveConfigObject(ctx, req.(*ConfigObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_UpdateConfigObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).UpdateConfigObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.config.v1.Config/UpdateConfigObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).UpdateConfigObjects(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_DeleteConfigObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).DeleteConfigObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.config.v1.Config/DeleteConfigObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).DeleteConfigObject(ctx, req.(*ConfigObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_GetLabelKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).GetLabelKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.config.v1.Config/GetLabelKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).GetLabelKeys(ctx, req.(*GetLabelKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_GetLogConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).GetLogConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.config.v1.Config/GetLogConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).GetLogConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_SaveLogConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLevels)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).SaveLogConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.config.v1.Config/SaveLogConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).SaveLogConfig(ctx, req.(*LogLevels))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_GetScriptAnnotations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScriptAnnotationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).GetScriptAnnotations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.config.v1.Config/GetScriptAnnotations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).GetScriptAnnotations(ctx, req.(*GetScriptAnnotationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Config_ServiceDesc is the grpc.ServiceDesc for Config service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Config_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.config.v1.Config",
	HandlerType: (*ConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfigObject",
			Handler:    _Config_GetConfigObject_Handler,
		},
		{
			MethodName: "CountConfigObjects",
			Handler:    _Config_CountConfigObjects_Handler,
		},
		{
			MethodName: "SaveConfigObject",
			Handler:    _Config_SaveConfigObject_Handler,
		},
		{
			MethodName: "UpdateConfigObjects",
			Handler:    _Config_UpdateConfigObjects_Handler,
		},
		{
			MethodName: "DeleteConfigObject",
			Handler:    _Config_DeleteConfigObject_Handler,
		},
		{
			MethodName: "GetLabelKeys",
			Handler:    _Config_GetLabelKeys_Handler,
		},
		{
			MethodName: "GetLogConfig",
			Handler:    _Config_GetLogConfig_Handler,
		},
		{
			MethodName: "SaveLogConfig",
			Handler:    _Config_SaveLogConfig_Handler,
		},
		{
			MethodName: "GetScriptAnnotations",
			Handler:    _Config_GetScriptAnnotations_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListConfigObjects",
			Handler:       _Config_ListConfigObjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "config/v1/config.proto",
}
