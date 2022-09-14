package database

import (
	"log"

	"github.com/bwoff11/frens/internal/models"
)

func DeleteStatusMediaByStatusID(id string) bool {
	if err := database.Where("status_id = ?", id).Delete(&models.StatusMedia{}).Error; err != nil {
		log.Println("Error deleting status media:", err)
		return false
	}
	return true
}
