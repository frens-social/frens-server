package models

type Reaction struct {
	ID        uint64 `json:"id"`
	AccountID uint64 `json:"account_id"`
	StatusID  uint64 `json:"status_id"`
	Emoji     string `json:"emoji"`
}
