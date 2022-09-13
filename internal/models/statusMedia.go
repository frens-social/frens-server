package models

import (
	"time"

	"github.com/google/uuid"
)

type StatusMedia struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	StatusID  uint64    `json:"status_id"`
	Status    Status    `json:"status" gorm:"-"`
	MediaID   uuid.UUID `json:"media_id"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
}
