package database

import "github.com/bwoff11/frens/internal/models"

// Partial implementation. Needs to take followed users into user.
func GetHomeFeed(continueFrom *int) ([]*models.Status, error) {
	limit := 20

	var statuses []*models.Status
	if err := database.
		Preload("User").
		Where("privacy = ?", "public").
		Where("draft = ?", false).
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
		Preload("User").
		Where("privacy = ?", "public").
		Where("draft = ?", false).
		Order("created_at desc").
		Limit(limit).
		Find(&statuses).Error; err != nil {
		return nil, err
	}
	return statuses, nil
}

func GetUserFeed(userId string, continueFrom *int) ([]*models.Status, error) {
	limit := 20

	var statuses []*models.Status
	if err := database.
		Preload("User").
		Where("privacy = ?", "public").
		Where("draft = ?", false).
		Where("user_id = ?", userId).
		Order("created_at desc").
		Limit(limit).
		Find(&statuses).Error; err != nil {
		return nil, err
	}
	return statuses, nil
}
