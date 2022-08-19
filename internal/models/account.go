package models

import "time"

type Account struct {
	ID        uint      `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
	//ActivatedAt   time.Time `gorm:"not null"`
	//DeactivatedAt time.Time `gorm:"not null"`
	Username string `json:"username"`
	Password string `json:"-"`
}
