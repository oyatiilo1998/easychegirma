// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package setting_service

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

// GisServiceClient is the client API for GisService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GisServiceClient interface {
	Create(ctx context.Context, in *Features, opts ...grpc.CallOption) (*CreatedResponse, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Features, error)
	GetAll(ctx context.Context, in *GetAllDataRequest, opts ...grpc.CallOption) (*GetAllDataResponse, error)
}

type gisServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGisServiceClient(cc grpc.ClientConnInterface) GisServiceClient {
	return &gisServiceClient{cc}
}

func (c *gisServiceClient) Create(ctx context.Context, in *Features, opts ...grpc.CallOption) (*CreatedResponse, error) {
	out := new(CreatedResponse)
	err := c.cc.Invoke(ctx, "/genproto.GisService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gisServiceClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Features, error) {
	out := new(Features)
	err := c.cc.Invoke(ctx, "/genproto.GisService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gisServiceClient) GetAll(ctx context.Context, in *GetAllDataRequest, opts ...grpc.CallOption) (*GetAllDataResponse, error) {
	out := new(GetAllDataResponse)
	err := c.cc.Invoke(ctx, "/genproto.GisService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GisServiceServer is the server API for GisService service.
// All implementations must embed UnimplementedGisServiceServer
// for forward compatibility
type GisServiceServer interface {
	Create(context.Context, *Features) (*CreatedResponse, error)
	Get(context.Context, *GetReq) (*Features, error)
	GetAll(context.Context, *GetAllDataRequest) (*GetAllDataResponse, error)
	mustEmbedUnimplementedGisServiceServer()
}

// UnimplementedGisServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGisServiceServer struct {
}

func (UnimplementedGisServiceServer) Create(context.Context, *Features) (*CreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGisServiceServer) Get(context.Context, *GetReq) (*Features, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGisServiceServer) GetAll(context.Context, *GetAllDataRequest) (*GetAllDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedGisServiceServer) mustEmbedUnimplementedGisServiceServer() {}

// UnsafeGisServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GisServiceServer will
// result in compilation errors.
type UnsafeGisServiceServer interface {
	mustEmbedUnimplementedGisServiceServer()
}

func RegisterGisServiceServer(s grpc.ServiceRegistrar, srv GisServiceServer) {
	s.RegisterService(&GisService_ServiceDesc, srv)
}

func _GisService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Features)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GisServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.GisService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GisServiceServer).Create(ctx, req.(*Features))
	}
	return interceptor(ctx, in, info, handler)
}

func _GisService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GisServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.GisService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GisServiceServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GisService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GisServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.GisService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GisServiceServer).GetAll(ctx, req.(*GetAllDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GisService_ServiceDesc is the grpc.ServiceDesc for GisService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GisService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.GisService",
	HandlerType: (*GisServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GisService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GisService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _GisService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gis_service.proto",
}
