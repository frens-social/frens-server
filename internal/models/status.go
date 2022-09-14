package models

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type StatusPrivacy string
type StatusState string

const (
	StatusPrivacyPublic    StatusPrivacy = "public"
	StatusPrivacyLocal     StatusPrivacy = "local"
	StatusPrivacyFollowers StatusPrivacy = "followers"
	StatusPrivacyPrivate   StatusPrivacy = "private"
)

type Status struct {
	ID             uint64        `json:"id" gorm:"primary_key"`
	CreatedAt      time.Time     `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time     `gorm:"not null" json:"updated_at"`
	UserID         uint64        `json:"-" gorm:"column:user_id"`
	User           User          `json:"user" gorm:"foreignkey:UserID"`
	Text           string        `json:"text" validate:"required"`
	Privacy        StatusPrivacy `json:"privacy" validate:"required"`
	Draft          bool          `json:"draft"`
	StatusMediaIDs uuid.UUID     `json:"-" gorm:"foreignkey:StatusID;constraint:OnDelete:CASCADE;"`
	StatusMedia    []StatusMedia `json:"media" gorm:"foreignkey:StatusID;constraint:OnDelete:CASCADE;"`
}

func (s *Status) Validate() error {
	return validator.New().Struct(s)
}
