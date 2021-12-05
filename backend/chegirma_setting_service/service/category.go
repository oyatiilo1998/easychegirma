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

type CategoryService struct {
	logger  logger.Logger
	storage storage.StorageI
	ss.UnimplementedCategoryServiceServer
}

func NewCategoryService(db *db.Database, log logger.Logger) *CategoryService {
	return &CategoryService{
		storage: storage.NewStorage(db),
		logger:  log,
	}
}

func (bs *CategoryService) Create(ctx context.Context, req *ss.Category) (*ss.CategoryCreatedResponse, error) {
	ID, err := bs.storage.Category().Create(req)

	if err != nil {
		bs.logger.Error("error while creating Category error -> ", logger.Error(err))
		return nil, helper.HandleError(bs.logger, err, "error while creating Category error -> ", req)
	}
	return &ss.CategoryCreatedResponse{
		Id: ID,
	}, nil
}

func (bs *CategoryService) Get(ctx context.Context, req *ss.CategoryGetReq) (*ss.Category, error) {

	response, err := bs.storage.Category().Get(req.Id)

	if err != nil {
		return nil, helper.HandleError(bs.logger, err, "error while getting Category error -> ", req)
	}
	return response, nil
}
func (bs *CategoryService) GetAll(ctx context.Context, req *ss.GetAllCategoriesRequest) (*ss.GetAllCategoriesResponse, error) {
	fmt.Println("----------->")
	response, count, err := bs.storage.Category().GetAll(req.Page, req.Limit, req.Soato, req.Name)

	if err != nil {
		return nil, helper.HandleError(bs.logger, err, "error while getting Category error -> ", req)
	}
	return &ss.GetAllCategoriesResponse{
		Count:      count,
		Categories: response,
	}, nil
}

func (bs *CategoryService) Update(ctx context.Context, req *ss.Category) (*empty.Empty, error) {
	if err := bs.storage.Category().Update(req); err != nil {
		bs.logger.Error("error while updating Category error -> ", logger.Error(err))
		return &empty.Empty{}, helper.HandleError(bs.logger, err, "error while updating Category error -> ", req)
	}
	return &emptypb.Empty{}, nil
}

func (bs *CategoryService) CategoryExists(ctx context.Context, req *ss.CCExistsRequest) (*ss.CCExistsResponse, error) {
	exist, err := bs.storage.Category().CategoryExists(req.Id)
	if err != nil {
		bs.logger.Error("error while checking if Category exist error -> ", logger.Error(err))
		return &ss.CCExistsResponse{Exist: false}, helper.HandleError(bs.logger, err, "error while checking if Category exist error -> ", req)
	}

	return &ss.CCExistsResponse{Exist: exist}, nil
}

func (bs *CategoryService) GetByCode(ctx context.Context, req *ss.GetCategoryByCodeRequest) (*ss.Category, error) {

	response, err := bs.storage.Category().GetByCode(int(req.Code))

	if err != nil {
		return nil, helper.HandleError(bs.logger, err, "error while getting Category error -> ", req)
	}
	return response, nil
}

// func (bs *CategoryService) Delete(ctx context.Context, req *ss.DeleteRequest) (*empty.Empty, error) {
// 	fmt.Println(req)
// 	if err := bs.storage.Category().Delete(req.Id); err != nil {
// 		bs.logger.Error("error while deleting Category error -> ", logger.Error(err))
// 		return &empty.Empty{}, helper.HandleError(bs.logger, err, "error while deleting Category error -> ", req)
// 	}
// 	return &empty.Empty{}, nil
// }
