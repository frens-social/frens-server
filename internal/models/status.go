package models

import (
	"time"

	"github.com/go-playground/validator"
)

type StatusPrivacy string
type StatusState string

const (
	PrivacyPublic  StatusPrivacy = "public"
	PrivacyLocal   StatusPrivacy = "local"
	PrivacyFriends StatusPrivacy = "friends"
	PrivacyPrivate StatusPrivacy = "private"
)

const (
	StateDraft     StatusState = "draft"
	StatePublished StatusState = "published"
)

type Status struct {
	ID        uint64        `json:"id" gorm:"primary_key"`
	CreatedAt time.Time     `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time     `gorm:"not null" json:"updated_at"`
	UserID    uint64        `json:"-" gorm:"column:user_id"`
	User      User          `json:"user" gorm:"foreignkey:UserID"`
	Text      string        `json:"text" validate:"required"`
	Privacy   StatusPrivacy `json:"privacy" validate:"required"`
	State     StatusState   `json:"state" validate:"required"`
}

func (s *Status) Validate() error {
	return validator.New().Struct(s)
}
