package models

import (
	"time"

	"github.com/google/uuid"
)

type StatusMedia struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key"`
	StatusID  uint      `json:"status_id"`
	Status    Status    `json:"status" gorm:"-"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
}
