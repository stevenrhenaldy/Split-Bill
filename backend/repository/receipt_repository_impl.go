package repository

import (
	"split-bill/backend/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReceiptRepositoryImpl struct {
	Db *gorm.DB
}

func NewReceiptRepositoryImpl(db *gorm.DB) ReceiptRepository {
	return &ReceiptRepositoryImpl{
		Db: db,
	}
}

// Delete implements ReceiptRepository.
func (r *ReceiptRepositoryImpl) Delete(id string) error {
	var receipt model.Receipt
	if err := r.Db.Where("id = ?", id).First(&receipt).Error; err != nil {
		return err
	}
	return r.Db.Delete(&receipt).Error
}

// FindAll implements ReceiptRepository.
func (r *ReceiptRepositoryImpl) FindAll() ([]*model.Receipt, error) {
	var receipts []*model.Receipt
	if err := r.Db.Find(&receipts).Error; err != nil {
		return nil, err
	}
	return receipts, nil
}

// FindByID implements ReceiptRepository.
func (r *ReceiptRepositoryImpl) FindByID(id string) (*model.Receipt, error) {
	var receipt model.Receipt
	if err := r.Db.Where("id = ?", id).First(&receipt).Error; err != nil {
		return nil, err
	}
	if receipt == (model.Receipt{}) {
		return nil, gorm.ErrRecordNotFound
	}
	return &receipt, nil
}

// Create implements ReceiptRepository.
func (r *ReceiptRepositoryImpl) Create(receipt *model.Receipt) (*model.Receipt, error) {
	// Generate a new UUID
	receipt.ID = uuid.New()

	// Set the created time and updated time
	receipt.UpdatedAt = time.Now()
	receipt.CreatedAt = time.Now()
	return receipt, r.Db.Create(receipt).Error
}

// Update implements ReceiptRepository.
func (r *ReceiptRepositoryImpl) Update(receipt *model.Receipt) (*model.Receipt, error) {
	// Set the updated time
	receipt.UpdatedAt = time.Now()
	if err := r.Db.Save(receipt).Error; err != nil {
		return nil, err
	}
	return receipt, nil
}
