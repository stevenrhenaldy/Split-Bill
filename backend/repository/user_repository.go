package repository

import (
	"split-bill/backend/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByUUID(uuid.UUID) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByUsernameOrEmail(input string) (*model.User, error)
	FindAll() ([]*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id string) error
}
