package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Share struct {
	ID             uuid.UUID       `gorm:"type:uuid"`
	Name           string          `gorm:"not null" json:"name"`
	ReceiptID      uuid.UUID       `gorm:"not null" json:"receipt_id"`
	Receipt        Receipt         `gorm:"foreignKey:ReceiptID" json:"receipt"`
	Status         int8            `gorm:"not null" json:"status"` // 0: pending, 1: completed, 2: cancelled
	UserID         *uuid.UUID      `gorm:"null" json:"user_id"`
	TaxAmount      decimal.Decimal `gorm:"type:decimal(16,4)" json:"tax_amount"`
	ServiceAmount  decimal.Decimal `gorm:"type:decimal(16,4)" json:"service_amount"`
	DiscountAmount decimal.Decimal `gorm:"type:decimal(16,4)" json:"discount_amount"`
}
