package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username          string    `gorm:"not null;uniqueIndex" json:"username"`
	Name              string    `gorm:"not null" json:"name"`
	Email             string    `gorm:"not null;uniqueIndex" json:"email"`
	Password          string    `gorm:"not null" json:"password"`
	DefaultCurrencyID string    `gorm:"type:varchar(3);not null" json:"default_currency_id"`
	DefaultCurrency   Currency  `gorm:"foreignKey:DefaultCurrencyID" json:"default_currency"`
	EmailVerifiedAt   time.Time `gorm:"default:current_timestamp" json:"email_verified_at"`
	CreatedAt         time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt         time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
