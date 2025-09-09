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
func (r *ReceiptRepositoryImpl) Delete(user *model.User, id string) error {
	var receipt model.Receipt
	if err := r.Db.Where("id = ? AND user_id = ?", id, user.ID).First(&receipt).Error; err != nil {
		return err
	}
	return r.Db.Delete(&receipt).Error
}

// FindAll implements ReceiptRepository.
func (r *ReceiptRepositoryImpl) FindAll(user *model.User) ([]*model.Receipt, error) {
	var receipts []*model.Receipt
	if err := r.Db.Where("user_id = ?", user.ID).Find(&receipts).Error; err != nil {
		return nil, err
	}
	return receipts, nil
}

// FindByID implements ReceiptRepository.
func (r *ReceiptRepositoryImpl) FindByID(user *model.User, id string) (*model.Receipt, error) {
	var receipt model.Receipt
	if err := r.Db.Where("id = ? AND user_id = ?", id, user.ID).First(&receipt).Error; err != nil {
		return nil, err
	}
	if receipt == (model.Receipt{}) {
		return nil, gorm.ErrRecordNotFound
	}
	return &receipt, nil
}

// Create implements ReceiptRepository.
func (r *ReceiptRepositoryImpl) Create(user *model.User, receipt *model.Receipt) (*model.Receipt, error) {
	// Generate a new UUID
	receipt.ID = uuid.New()

	// Set the created time and updated time
	receipt.UserID = user.ID
	receipt.UpdatedAt = time.Now()
	receipt.CreatedAt = time.Now()
	return receipt, r.Db.Create(receipt).Error
}

// Update implements ReceiptRepository.
func (r *ReceiptRepositoryImpl) Update(user *model.User, receipt *model.Receipt) (*model.Receipt, error) {
	// Ensure the receipt belongs to the user
	var existingReceipt model.Receipt
	if err := r.Db.Where("id = ? AND user_id = ?", receipt.ID, user.ID).First(&existingReceipt).Error; err != nil {
		return nil, err
	}
	// Update fields
	existingReceipt.Name = receipt.Name
	existingReceipt.Description = receipt.Description
	// Set the updated time
	existingReceipt.UpdatedAt = time.Now()
	if err := r.Db.Save(&existingReceipt).Error; err != nil {
		return nil, err
	}
	return &existingReceipt, nil
}
