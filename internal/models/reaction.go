package models

type Reaction struct {
	ID       uint64 `json:"id"`
	UserID   uint64 `json:"user_id"`
	StatusID uint64 `json:"status_id"`
	Emoji    string `json:"emoji"`
}
