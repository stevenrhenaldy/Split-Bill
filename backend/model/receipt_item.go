package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ReceiptItem struct {
	ID                uuid.UUID       `gorm:"type:uuid"`
	ReceiptID         uuid.UUID       `gorm:"not null" json:"receipt_id"`
	Receipt           Receipt         `gorm:"foreignKey:ReceiptID" json:"receipt"`
	DisplayName       string          `gorm:"not null" json:"display_name"`
	Quantity          int             `gorm:"not null" json:"quantity"`
	TotalItemPrice    decimal.Decimal `gorm:"type:decimal(16,4)" json:"total_item_price"`
	TotalItemDiscount decimal.Decimal `gorm:"type:decimal(16,4)" json:"total_item_discount"`
}
