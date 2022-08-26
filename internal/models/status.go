package models

import "github.com/go-playground/validator"

type Status struct {
	ID        uint64  `json:"id" gorm:"primary_key"`
	AccountID uint64  `json:"-" gorm:"column:account_id"`
	Account   Account `json:"account" gorm:"foreignkey:AccountID"`
	Text      string  `json:"text" validate:"required"`
}

func (s *Status) Validate() error {
	return validator.New().Struct(s)
}
