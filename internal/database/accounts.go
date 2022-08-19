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

func GetAccount(id string, account *models.Account) bool {
	if err := database.
		Where("id = ?", id).
		First(account).
		Error; err != nil {
		return false
	}
	return true
}
