package service

import (
	"context"
	"fmt"

	ss "chegirma_setting_service/genproto/setting_service"
	"chegirma_setting_service/pkg/helper"
	"chegirma_setting_service/pkg/logger"
	"chegirma_setting_service/storage"

	"github.com/golang/protobuf/ptypes/empty"
	db "go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductService struct {
	logger  logger.Logger
	storage storage.StorageI
	ss.UnimplementedProductServiceServer
}

func NewProductService(db *db.Database, log logger.Logger) *ProductService {
	return &ProductService{
		storage: storage.NewStorage(db),
		logger:  log,
	}
}

func (bs *ProductService) Create(ctx context.Context, req *ss.CreateUpdateProduct) (*ss.CreatedResponse, error) {
	ID, err := bs.storage.Product().Create(req)

	if err != nil {
		bs.logger.Error("error while creating Product error -> ", logger.Error(err))
		return nil, helper.HandleError(bs.logger, err, "error while creating Product error -> ", req)
	}
	return &ss.CreatedResponse{
		Id: ID,
	}, nil
}

func (bs *ProductService) Get(ctx context.Context, req *ss.GetReq) (*ss.Product, error) {

	response, err := bs.storage.Product().Get(req.Id)

	if err != nil {
		return nil, helper.HandleError(bs.logger, err, "error while getting Product error -> ", req)
	}
	return response, nil
}
func (bs *ProductService) GetAll(ctx context.Context, req *ss.GetAllProductsRequest) (*ss.GetAllProductsResponse, error) {
	fmt.Println("----------->")
	response, count, err := bs.storage.Product().GetAll(req.Page, req.Limit, req.CategoryId, req.OwnerId, req.Name)

	if err != nil {
		return nil, helper.HandleError(bs.logger, err, "error while getting Product error -> ", req)
	}
	return &ss.GetAllProductsResponse{
		Count:    count,
		Products: response,
	}, nil
}

func (bs *ProductService) Update(ctx context.Context, req *ss.CreateUpdateProduct) (*empty.Empty, error) {
	if err := bs.storage.Product().Update(req); err != nil {
		bs.logger.Error("error while updating Product error -> ", logger.Error(err))
		return &empty.Empty{}, helper.HandleError(bs.logger, err, "error while updating Product error -> ", req)
	}
	return &emptypb.Empty{}, nil
}

func (bs *ProductService) ProductExists(ctx context.Context, req *ss.SSExistsRequest) (*ss.SSExistsResponse, error) {
	exist, err := bs.storage.Product().ProductExists(req.Id)
	if err != nil {
		bs.logger.Error("error while checking if Product exist error -> ", logger.Error(err))
		return &ss.SSExistsResponse{Exist: false}, helper.HandleError(bs.logger, err, "error while checking if Product exist error -> ", req)
	}

	return &ss.SSExistsResponse{Exist: exist}, nil
}

func (bs *ProductService) Delete(ctx context.Context, req *ss.DeleteProductRequest) (*empty.Empty, error) {
	fmt.Println(req)
	if err := bs.storage.Product().Delete(req.Id); err != nil {
		bs.logger.Error("error while deleting Product error -> ", logger.Error(err))
		return &empty.Empty{}, helper.HandleError(bs.logger, err, "error while deleting Product error -> ", req)
	}
	return &empty.Empty{}, nil
}
