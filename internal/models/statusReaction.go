package models

type StatusReaction struct {
	ID       uint64 `json:"id"`
	UserID   uint64 `json:"user_id"`
	User     User   `json:"-"`
	StatusID uint64 `json:"status_id"`
	Status   Status `json:"-"`
	Emoji    string `json:"emoji"`
}
