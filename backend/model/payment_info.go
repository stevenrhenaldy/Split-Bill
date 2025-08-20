package model

import (
	"github.com/google/uuid"
)

type PaymentInfo struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID       uuid.UUID `gorm:"not null" json:"user_id"`
	User         User      `gorm:"foreignKey:UserID" json:"user"`
	Description  string    `gorm:"" json:"description"`
	BankAccount  string    `gorm:"" json:"bank_account"`
	FileLocation string    `gorm:"" json:"file_location"`
}
