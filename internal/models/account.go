package models

import "time"

type Account struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
	//ActivatedAt   time.Time `gorm:"not null"`
	//DeactivatedAt time.Time `gorm:"not null"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Password    string `json:"-"`
	AvatarURL   string `json:"avatar_url"`
	BannerURL   string `json:"banner_url"`
}
