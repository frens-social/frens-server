package database

import "github.com/bwoff11/frens/internal/models"

func CreateReaction(reaction *models.Reaction) error {
	if err := database.Create(reaction).Error; err != nil {
		return err
	}
	return nil
}

func GetStatusReactions(statusID uint64) ([]models.Reaction, error) {
	var reactions []models.Reaction
	if err := database.Where("status_id = ?", statusID).Find(&reactions).Error; err != nil {
		return nil, err
	}
	return reactions, nil
}
