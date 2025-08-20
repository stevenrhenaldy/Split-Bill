package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ReceiptItemShare struct {
	ID            uuid.UUID       `gorm:"type:uuid"`
	ReceiptItemID uuid.UUID       `gorm:"not null" json:"receipt_item_id"`
	ReceiptItem   ReceiptItem     `gorm:"foreignKey:ReceiptItemID" json:"receipt_item"`
	ShareID       uuid.UUID       `gorm:"not null" json:"share_id"`
	Share         Share           `gorm:"foreignKey:ShareID" json:"share"`
	ShareType     int8            `gorm:"not null" json:"share_type"`
	ShareQuantity float64         `gorm:"not null" json:"share_quantity"`
	ShareAmount   decimal.Decimal `gorm:"type:decimal(16,4)" json:"share_amount"`
}
