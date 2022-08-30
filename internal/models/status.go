package models

import (
	"time"

	"github.com/go-playground/validator"
)

type StatusPrivacy string

const (
	PrivacyPublic  StatusPrivacy = "public"
	PrivacyLocal   StatusPrivacy = "local"
	PrivacyFriends StatusPrivacy = "friends"
	PrivacyPrivate StatusPrivacy = "private"
)

type Status struct {
	ID        uint64        `json:"id" gorm:"primary_key"`
	CreatedAt time.Time     `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time     `gorm:"not null" json:"updated_at"`
	AccountID uint64        `json:"-" gorm:"column:account_id"`
	Account   Account       `json:"account" gorm:"foreignkey:AccountID"`
	Text      string        `json:"text" validate:"required"`
	Privacy   StatusPrivacy `json:"privacy" validate:"required"`
}

func (s *Status) Validate() error {
	return validator.New().Struct(s)
}
