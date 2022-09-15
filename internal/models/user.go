package models

import "time"

type UserType string

const (
	UserTypePerson UserType = "person"
	UserTypeBot    UserType = "bot"
)

type User struct {
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

	//ActivityPub
	Inbox     string `json:"inbox"`
	Outbox    string `json:"outbox"`
	Following string `json:"following"`
	Followers string `json:"followers"`
	Liked     string `json:"liked"`
}
