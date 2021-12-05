// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package setting_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CityServiceClient is the client API for CityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CityServiceClient interface {
	Create(ctx context.Context, in *City, opts ...grpc.CallOption) (*CreatedResponse, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*City, error)
	GetAll(ctx context.Context, in *GetAllCitiesRequest, opts ...grpc.CallOption) (*GetAllCitiesResponse, error)
	Update(ctx context.Context, in *City, opts ...grpc.CallOption) (*empty.Empty, error)
	CityExists(ctx context.Context, in *SSExistsRequest, opts ...grpc.CallOption) (*SSExistsResponse, error)
	GetByCode(ctx context.Context, in *GetByCodeRequest, opts ...grpc.CallOption) (*City, error)
}

type cityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCityServiceClient(cc grpc.ClientConnInterface) CityServiceClient {
	return &cityServiceClient{cc}
}

func (c *cityServiceClient) Create(ctx context.Context, in *City, opts ...grpc.CallOption) (*CreatedResponse, error) {
	out := new(CreatedResponse)
	err := c.cc.Invoke(ctx, "/genproto.CityService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*City, error) {
	out := new(City)
	err := c.cc.Invoke(ctx, "/genproto.CityService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) GetAll(ctx context.Context, in *GetAllCitiesRequest, opts ...grpc.CallOption) (*GetAllCitiesResponse, error) {
	out := new(GetAllCitiesResponse)
	err := c.cc.Invoke(ctx, "/genproto.CityService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) Update(ctx context.Context, in *City, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.CityService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) CityExists(ctx context.Context, in *SSExistsRequest, opts ...grpc.CallOption) (*SSExistsResponse, error) {
	out := new(SSExistsResponse)
	err := c.cc.Invoke(ctx, "/genproto.CityService/CityExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) GetByCode(ctx context.Context, in *GetByCodeRequest, opts ...grpc.CallOption) (*City, error) {
	out := new(City)
	err := c.cc.Invoke(ctx, "/genproto.CityService/GetByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CityServiceServer is the server API for CityService service.
// All implementations must embed UnimplementedCityServiceServer
// for forward compatibility
type CityServiceServer interface {
	Create(context.Context, *City) (*CreatedResponse, error)
	Get(context.Context, *GetReq) (*City, error)
	GetAll(context.Context, *GetAllCitiesRequest) (*GetAllCitiesResponse, error)
	Update(context.Context, *City) (*empty.Empty, error)
	CityExists(context.Context, *SSExistsRequest) (*SSExistsResponse, error)
	GetByCode(context.Context, *GetByCodeRequest) (*City, error)
	mustEmbedUnimplementedCityServiceServer()
}

// UnimplementedCityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCityServiceServer struct {
}

func (UnimplementedCityServiceServer) Create(context.Context, *City) (*CreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCityServiceServer) Get(context.Context, *GetReq) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCityServiceServer) GetAll(context.Context, *GetAllCitiesRequest) (*GetAllCitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCityServiceServer) Update(context.Context, *City) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCityServiceServer) CityExists(context.Context, *SSExistsRequest) (*SSExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CityExists not implemented")
}
func (UnimplementedCityServiceServer) GetByCode(context.Context, *GetByCodeRequest) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCode not implemented")
}
func (UnimplementedCityServiceServer) mustEmbedUnimplementedCityServiceServer() {}

// UnsafeCityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CityServiceServer will
// result in compilation errors.
type UnsafeCityServiceServer interface {
	mustEmbedUnimplementedCityServiceServer()
}

func RegisterCityServiceServer(s grpc.ServiceRegistrar, srv CityServiceServer) {
	s.RegisterService(&CityService_ServiceDesc, srv)
}

func _CityService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(City)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CityService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).Create(ctx, req.(*City))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CityService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CityService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).GetAll(ctx, req.(*GetAllCitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(City)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CityService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).Update(ctx, req.(*City))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_CityExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).CityExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CityService/CityExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).CityExists(ctx, req.(*SSExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_GetByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).GetByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CityService/GetByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).GetByCode(ctx, req.(*GetByCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CityService_ServiceDesc is the grpc.ServiceDesc for CityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.CityService",
	HandlerType: (*CityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CityService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CityService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CityService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CityService_Update_Handler,
		},
		{
			MethodName: "CityExists",
			Handler:    _CityService_CityExists_Handler,
		},
		{
			MethodName: "GetByCode",
			Handler:    _CityService_GetByCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "city_service.proto",
}
