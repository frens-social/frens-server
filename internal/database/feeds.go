package database

import "github.com/bwoff11/frens/internal/models"

// Partial implementation. Needs to take followed accounts into account.
func GetHomeFeed(continueFrom *int) ([]*models.Status, error) {
	limit := 20

	var statuses []*models.Status
	if err := database.
		Preload("Account").
		Where("privacy = ?", "public").
		Order("created_at desc").
		Limit(limit).
		Find(&statuses).Error; err != nil {
		return nil, err
	}
	return statuses, nil
}

func GetPublicFeed(continueFrom *int) ([]*models.Status, error) {
	limit := 20

	var statuses []*models.Status
	if err := database.
		Preload("Account").
		Where("privacy = ?", "public").
		Order("created_at desc").
		Limit(limit).
		Find(&statuses).Error; err != nil {
		return nil, err
	}
	return statuses, nil
}
