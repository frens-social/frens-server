package models

import "time"

type Follow struct {
	ID              uint64    `json:"id" gorm:"primary_key"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
	SourceAccountID uint64    `json:"-" gorm:"column:source_account_id"`
	TargetAccountID uint64    `json:"-" gorm:"column:target_account_id"`
}
