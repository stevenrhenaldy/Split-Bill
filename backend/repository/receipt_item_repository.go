package repository

import "split-bill/backend/model"

type ReceiptItemRepository interface {
	Create(*model.Receipt, *model.ReceiptItem) (*model.ReceiptItem, error)
	FindByID(*model.Receipt, string) (*model.ReceiptItem, error)
	FindAll(*model.Receipt) ([]*model.ReceiptItem, error)
	Update(*model.Receipt, *model.ReceiptItem) (*model.ReceiptItem, error)
	Delete(*model.Receipt, string) error
}
