package storage

import (
	"chegirma_setting_service/storage/mongo"
	"chegirma_setting_service/storage/repo"

	db "go.mongodb.org/mongo-driver/mongo"
)

type StorageI interface {
	Product() repo.ProductStorageI
	Category() repo.CategoryStorageI
	User() repo.UserStorageI
}

type storageMDB struct {
	categoryRepo repo.CategoryStorageI
	productRepo  repo.ProductStorageI
	userRepo     repo.UserStorageI
}

func NewStorage(db *db.Database) StorageI {
	return &storageMDB{
		categoryRepo: mongo.NewCategoryRepo(db),
		productRepo:  mongo.NewProductRepo(db),
		userRepo:     mongo.NewUserRepo(db),
	}
}

func (s *storageMDB) Category() repo.CategoryStorageI {
	return s.categoryRepo
}

func (s *storageMDB) Product() repo.ProductStorageI {
	return s.productRepo
}

func (s *storageMDB) User() repo.UserStorageI {
	return s.userRepo
}
