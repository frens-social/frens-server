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

func GetStatuses(count int) ([]models.Status, error) {
	var statuses []models.Status
	if err := database.Limit(count).Find(&statuses).Error; err != nil {
		return nil, err
	}
	return statuses, nil
}
