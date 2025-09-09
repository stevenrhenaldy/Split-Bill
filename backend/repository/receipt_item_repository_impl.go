package repository

import (
	"split-bill/backend/model"

	"gorm.io/gorm"
)

type ReceiptItemRepositoryImpl struct {
	Db *gorm.DB
}

func NewReceiptItemRepositoryImpl(db *gorm.DB) ReceiptItemRepository {
	return &ReceiptItemRepositoryImpl{
		Db: db,
	}
}

// Create implements ReceiptItemRepository.
func (r *ReceiptItemRepositoryImpl) Create(*model.Receipt, *model.ReceiptItem) (*model.ReceiptItem, error) {
	panic("unimplemented")
}

// Delete implements ReceiptItemRepository.
func (r *ReceiptItemRepositoryImpl) Delete(*model.Receipt, string) error {
	panic("unimplemented")
}

// FindAll implements ReceiptItemRepository.
func (r *ReceiptItemRepositoryImpl) FindAll(*model.Receipt) ([]*model.ReceiptItem, error) {
	panic("unimplemented")
}

// FindByID implements ReceiptItemRepository.
func (r *ReceiptItemRepositoryImpl) FindByID(*model.Receipt, string) (*model.ReceiptItem, error) {
	panic("unimplemented")
}

// Update implements ReceiptItemRepository.
func (r *ReceiptItemRepositoryImpl) Update(*model.Receipt, *model.ReceiptItem) (*model.ReceiptItem, error) {
	panic("unimplemented")
}
