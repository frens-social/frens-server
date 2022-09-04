package models

import "time"

type Follow struct {
	ID           uint64    `json:"id" gorm:"primary_key"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	SourceUserID uint64    `json:"-" gorm:"column:source_user_id"`
	TargetUserID uint64    `json:"-" gorm:"column:target_user_id"`
}
