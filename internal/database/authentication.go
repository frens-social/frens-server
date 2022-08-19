package database

import "github.com/bwoff11/frens/internal/models"

func Authenticate(username string, password string) bool {
	var account models.Account
	if err := database.
		Where("username = ?", username).
		First(&account).
		Error; err != nil {
		return false
	}
	return true
}
