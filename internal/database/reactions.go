package database

import "github.com/bwoff11/frens/internal/models"

func CreateReaction(reaction *models.StatusReaction) error {
	if err := database.
		Create(reaction).
		Preload("User").
		Preload("Status").
		Error; err != nil {
		return err
	}
	return nil
}

func GetStatusReactions(statusID uint64) ([]models.StatusReaction, error) {
	var reactions []models.StatusReaction
	if err := database.
		Where("status_id = ?", statusID).
		Preload("User").
		Preload("Status").
		Find(&reactions).Error; err != nil {
		return nil, err
	}
	return reactions, nil
}

func GetUserStatusReactions(statusID uint64, userID uint64) ([]models.StatusReaction, error) {
	var reactions []models.StatusReaction
	if err := database.
		Where("status_id = ? AND user_id = ?", statusID, userID).
		Preload("User").
		Preload("Status").
		Find(&reactions).Error; err != nil {
		return nil, err
	}
	return reactions, nil
}
