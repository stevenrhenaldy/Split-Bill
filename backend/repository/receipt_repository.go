package repository

import "split-bill/backend/model"

type ReceiptRepository interface {
	Create(*model.User, *model.Receipt) (*model.Receipt, error)
	FindByID(*model.User, string) (*model.Receipt, error)
	FindAll(*model.User) ([]*model.Receipt, error)
	Update(*model.User, *model.Receipt) (*model.Receipt, error)
	Delete(*model.User, string) error
}
