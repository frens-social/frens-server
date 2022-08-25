package models

import "github.com/go-playground/validator"

type Status struct {
	ID   uint   `json:"-" gorm:"primary_key"`
	Text string `json:"text" validate:"required"`
}

func (s *Status) Validate() error {
	return validator.New().Struct(s)
}
