package repo

import (
	ss "chegirma_setting_service/genproto/setting_service"
)

type ProductStorageI interface {
	Create(city *ss.CreateUpdateProduct) (string, error)
	Get(id string) (*ss.Product, error)
	GetAll(page, limit uint32, category_id, owner_id, name string) ([]*ss.Product, uint32, error)
	Update(city *ss.CreateUpdateProduct) error
	Delete(id string) error
	ProductExists(id string) (bool, error)
}
