package repo

import (
	ss "chegirma_setting_service/genproto/setting_service"
)

type CategoryStorageI interface {
	Create(city *ss.Category) (string, error)
	Get(id string) (*ss.Category, error)
	GetAll(page, limit, soato uint32, name string) ([]*ss.Category, uint32, error)
	Update(city *ss.Category) error
	Delete(id string) error
	CategoryExists(id string) (bool, error)
	GetByCode(soato int) (*ss.Category, error)
}
