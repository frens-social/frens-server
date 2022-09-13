package database

import "github.com/bwoff11/frens/internal/models"

func CreateStatusMedia(statusMedia *models.StatusMedia) bool {
	if err := database.Create(&statusMedia).Error; err != nil {
		return false
	}
	return true
}
