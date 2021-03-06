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

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	Create(ctx context.Context, in *Category, opts ...grpc.CallOption) (*CategoryCreatedResponse, error)
	Get(ctx context.Context, in *CategoryGetReq, opts ...grpc.CallOption) (*Category, error)
	GetAll(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error)
	Update(ctx context.Context, in *Category, opts ...grpc.CallOption) (*empty.Empty, error)
	CategoryExists(ctx context.Context, in *CCExistsRequest, opts ...grpc.CallOption) (*CCExistsResponse, error)
	GetByCode(ctx context.Context, in *GetCategoryByCodeRequest, opts ...grpc.CallOption) (*Category, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) Create(ctx context.Context, in *Category, opts ...grpc.CallOption) (*CategoryCreatedResponse, error) {
	out := new(CategoryCreatedResponse)
	err := c.cc.Invoke(ctx, "/genproto.CategoryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Get(ctx context.Context, in *CategoryGetReq, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/genproto.CategoryService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetAll(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error) {
	out := new(GetAllCategoriesResponse)
	err := c.cc.Invoke(ctx, "/genproto.CategoryService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Update(ctx context.Context, in *Category, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.CategoryService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CategoryExists(ctx context.Context, in *CCExistsRequest, opts ...grpc.CallOption) (*CCExistsResponse, error) {
	out := new(CCExistsResponse)
	err := c.cc.Invoke(ctx, "/genproto.CategoryService/CategoryExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetByCode(ctx context.Context, in *GetCategoryByCodeRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/genproto.CategoryService/GetByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	Create(context.Context, *Category) (*CategoryCreatedResponse, error)
	Get(context.Context, *CategoryGetReq) (*Category, error)
	GetAll(context.Context, *GetAllCategoriesRequest) (*GetAllCategoriesResponse, error)
	Update(context.Context, *Category) (*empty.Empty, error)
	CategoryExists(context.Context, *CCExistsRequest) (*CCExistsResponse, error)
	GetByCode(context.Context, *GetCategoryByCodeRequest) (*Category, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) Create(context.Context, *Category) (*CategoryCreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCategoryServiceServer) Get(context.Context, *CategoryGetReq) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCategoryServiceServer) GetAll(context.Context, *GetAllCategoriesRequest) (*GetAllCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCategoryServiceServer) Update(context.Context, *Category) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCategoryServiceServer) CategoryExists(context.Context, *CCExistsRequest) (*CCExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryExists not implemented")
}
func (UnimplementedCategoryServiceServer) GetByCode(context.Context, *GetCategoryByCodeRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCode not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CategoryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Create(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CategoryService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Get(ctx, req.(*CategoryGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CategoryService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetAll(ctx, req.(*GetAllCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CategoryService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Update(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CategoryExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CCExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CategoryExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CategoryService/CategoryExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CategoryExists(ctx, req.(*CCExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.CategoryService/GetByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetByCode(ctx, req.(*GetCategoryByCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CategoryService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CategoryService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CategoryService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CategoryService_Update_Handler,
		},
		{
			MethodName: "CategoryExists",
			Handler:    _CategoryService_CategoryExists_Handler,
		},
		{
			MethodName: "GetByCode",
			Handler:    _CategoryService_GetByCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category_service.proto",
}
