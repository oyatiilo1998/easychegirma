package services

import (
	"fmt"

	"chegirma_api_gateway/config"
	ss "chegirma_api_gateway/genproto/setting_service"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	GlobalSettingService() ss.GlobalSettingServiceClient
	CategoryService() ss.CategoryServiceClient
	ProductService() ss.ProductServiceClient
	UserService() ss.UserServiceClient
}

type grpcClients struct {
	categoryService ss.CategoryServiceClient
	productService  ss.ProductServiceClient
	userService     ss.UserServiceClient
}

func (g grpcClients) CategoryService() ss.CategoryServiceClient {
	return g.categoryService
}
func (g grpcClients) ProductService() ss.ProductServiceClient {
	return g.productService
}
func (g grpcClients) UserService() ss.UserServiceClient {
	return g.userService
}

func NewGrpcClients(cfg config.Config) (ServiceManager, error) {
	connSettingService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.SettingServiceHost, cfg.SettingServicePort),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10*1024*1024)))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &grpcClients{
		categoryService: ss.NewCategoryServiceClient(connSettingService),
		productService:  ss.NewProductServiceClient(connSettingService),
		userService:     ss.NewUserServiceClient(connSettingService),
	}, nil
}
