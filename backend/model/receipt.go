package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Receipt struct {
	ID             uuid.UUID       `gorm:"type:uuid"`
	UserID         uuid.UUID       `gorm:"" json:"user_id"`
	User           User            `gorm:"" json:"user"`
	CurrencyID     string          `gorm:"type:varchar(3)" json:"currency_id"`
	Name           string          `gorm:"not null" json:"name"`
	Description    string          `gorm:"" json:"description"`
	PaymentInfoID  uuid.UUID       `gorm:"" json:"payment_info_id"`
	FileLocation   string          `gorm:"" json:"file_location"`
	TaxAmount      decimal.Decimal `gorm:"type:decimal(16,4)" json:"tax_amount"`
	ServiceAmount  decimal.Decimal `gorm:"type:decimal(16,4)" json:"service_amount"`
	DiscountAmount decimal.Decimal `gorm:"type:decimal(16,4)" json:"discount_amount"`
	TotalAmount    decimal.Decimal `gorm:"type:decimal(16,4)" json:"total_amount"`
	CreatedAt      time.Time       `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"default:current_timestamp" json:"updated_at"`
}
