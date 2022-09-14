package models

import (
	"time"

	"github.com/google/uuid"
)

type StatusMedia struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	StatusID  uint64    `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	Status    Status    `json:"status"`
}
