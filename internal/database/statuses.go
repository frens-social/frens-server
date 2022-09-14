package database

import (
	"github.com/bwoff11/frens/internal/models"
)

func CreateStatus(status *models.Status) bool {
	if err := database.Create(&status).Error; err != nil {
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

func GetStatus(id string) (*models.Status, error) {
	var status models.Status
	if err := database.
		Preload("User").
		Preload("StatusMedia").
		Where("id = ?", id).First(&status).Error; err != nil {
		return nil, err
	}

	return &status, nil
}

func UpdateStatus(status *models.Status) bool {
	if err := database.Save(status).Error; err != nil {
		return false
	}
	return true
}
