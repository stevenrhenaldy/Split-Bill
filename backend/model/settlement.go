package model

import (
	"time"

	"github.com/google/uuid"
)

type Settlement struct {
	ID                uuid.UUID `gorm:"type:uuid"`
	ShareID           uuid.UUID `gorm:"not null" json:"share_id"`
	Share             Share     `gorm:"foreignKey:ShareID" json:"share"`
	ImageFileLocation string    `gorm:"" json:"image_file_location"`
	CreatedAt         time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt         time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
