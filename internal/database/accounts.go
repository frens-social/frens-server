package database

import "github.com/bwoff11/frens/internal/models"

func CreateAccount(account *models.Account) bool {
	if err := database.
		Create(account).
		Error; err != nil {
		return false
	}
	return true
}

func GetAccount(id uint64) *models.Account {
	var account models.Account
	if err := database.
		Where("id = ?", id).
		First(&account).
		Error; err != nil {
		return nil
	}
	return &account
}
