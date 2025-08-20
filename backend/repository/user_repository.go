package repository

import "split-bill/backend/model"

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByID(id string) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindAll() ([]*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id string) error
}
