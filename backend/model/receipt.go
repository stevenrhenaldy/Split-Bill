package model

import (
	"time"

	"github.com/google/uuid"
)

type Receipt struct {
	ID          uuid.UUID `gorm:"type:uuid"`
	UserID      uuid.UUID `gorm:"" json:"user_id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"" json:"description"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
