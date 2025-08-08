package repository

import "split-bill/backend/model"

type ReceiptRepository interface {
	Create(receipt *model.Receipt) (*model.Receipt, error)
	FindByID(id string) (*model.Receipt, error)
	FindAll() ([]*model.Receipt, error)
	Update(receipt *model.Receipt) (*model.Receipt, error)
	Delete(id string) error
}
