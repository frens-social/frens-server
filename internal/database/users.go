package database

import "github.com/bwoff11/frens/internal/models"

func CreateUser(User *models.User) bool {
	if err := database.
		Create(User).
		Error; err != nil {
		return false
	}
	return true
}

func GetUser(id *uint64) *models.User {
	var User models.User
	if err := database.
		Where("id = ?", id).
		First(&User).
		Error; err != nil {
		return nil
	}
	return &User
}

func GetFollowedUserIDs(UserID *uint64) (*[]uint64, error) {
	var followedUserIDs []uint64
	if err := database.
		Table("follows").
		Select("followed_User_id").
		Where("User_id = ?", UserID).
		Pluck("followed_User_id", &followedUserIDs).
		Error; err != nil {
		return nil, err
	}
	return &followedUserIDs, nil
}
