package repo

import (
	ss "chegirma_setting_service/genproto/setting_service"
)

type UserStorageI interface {
	Create(city *ss.User) (string, error)
	Get(id string) (*ss.User, error)
	GetAll(page, limit uint32, name string) ([]*ss.User, uint32, error)
	Update(city *ss.User) error
	UserExists(login, password string) (*ss.UserExistsResponse, error)
}
