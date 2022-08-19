package database

import "github.com/bwoff11/frens/internal/models"

func CreateStatus(status *models.Status) bool {
	if err := database.Create(status).Error; err != nil {
		return false
	}
	return true
}

func DeleteStatus(id string) bool {
	if err := database.Delete(models.Status{}, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}
