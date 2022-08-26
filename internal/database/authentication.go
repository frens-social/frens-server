package database

import "github.com/bwoff11/frens/internal/models"

func Authenticate(username string, password string) *models.Account {
	var account models.Account
	if err := database.
		Where("username = ? AND password = ?", username, password).
		First(&account).
		Error; err != nil {
		return nil
	}
	return &account
}
