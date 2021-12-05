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

type UserService struct {
	logger  logger.Logger
	storage storage.StorageI
	ss.UnimplementedUserServiceServer
}

func NewUserService(db *db.Database, log logger.Logger) *UserService {
	return &UserService{
		storage: storage.NewStorage(db),
		logger:  log,
	}
}

func (bs *UserService) Create(ctx context.Context, req *ss.User) (*ss.UserCreatedResponse, error) {
	ID, err := bs.storage.User().Create(req)

	if err != nil {
		bs.logger.Error("error while creating User error -> ", logger.Error(err))
		return nil, helper.HandleError(bs.logger, err, "error while creating User error -> ", req)
	}
	return &ss.UserCreatedResponse{
		Id: ID,
	}, nil
}

func (bs *UserService) Get(ctx context.Context, req *ss.UserGetReq) (*ss.User, error) {

	response, err := bs.storage.User().Get(req.Id)

	if err != nil {
		return nil, helper.HandleError(bs.logger, err, "error while getting User error -> ", req)
	}
	return response, nil
}
func (bs *UserService) GetAll(ctx context.Context, req *ss.GetAllUsersRequest) (*ss.GetAllUsersResponse, error) {
	fmt.Println("----------->")
	response, count, err := bs.storage.User().GetAll(req.Page, req.Limit, req.Name)

	if err != nil {
		return nil, helper.HandleError(bs.logger, err, "error while getting User error -> ", req)
	}
	return &ss.GetAllUsersResponse{
		Count: count,
		Users: response,
	}, nil
}

func (bs *UserService) Update(ctx context.Context, req *ss.User) (*empty.Empty, error) {
	if err := bs.storage.User().Update(req); err != nil {
		bs.logger.Error("error while updating User error -> ", logger.Error(err))
		return &empty.Empty{}, helper.HandleError(bs.logger, err, "error while updating User error -> ", req)
	}
	return &emptypb.Empty{}, nil
}

func (bs *UserService) UserExists(ctx context.Context, req *ss.UserExistsRequest) (*ss.UserExistsResponse, error) {
	user, err := bs.storage.User().UserExists(req.Login, req.Password)
	if err != nil {
		bs.logger.Error("error while checking if User exist error -> ", logger.Error(err))
		return &ss.UserExistsResponse{Exist: false}, helper.HandleError(bs.logger, err, "error while checking if User exist error -> ", req)
	}

	return &ss.UserExistsResponse{Exist: user.Exist, Id: user.Id}, nil
}

// func (bs *UserService) Delete(ctx context.Context, req *ss.DeleteRequest) (*empty.Empty, error) {
// 	fmt.Println(req)
// 	if err := bs.storage.User().Delete(req.Id); err != nil {
// 		bs.logger.Error("error while deleting User error -> ", logger.Error(err))
// 		return &empty.Empty{}, helper.HandleError(bs.logger, err, "error while deleting User error -> ", req)
// 	}
// 	return &empty.Empty{}, nil
// }
