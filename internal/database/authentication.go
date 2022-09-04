package database

import "github.com/bwoff11/frens/internal/models"

func Authenticate(username string, password string) *models.User {
	var user models.User
	if err := database.
		Where("username = ? AND password = ?", username, password).
		First(&user).
		Error; err != nil {
		return nil
	}
	return &user
}
