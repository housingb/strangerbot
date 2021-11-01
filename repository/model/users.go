package model

import (
	"database/sql"
)

type User struct {
	ID            int64         `db:"id"`
	ChatID        int64         `db:"chat_id"`
	Available     bool          `db:"available"`
	MatchChatID   sql.NullInt64 `db:"match_chat_id"`
	PreviousMatch sql.NullInt64 `db:"previous_match"`
	AllowPictures bool          `db:"allow_pictures"`
}

func (p *User) TableName() string {
	return "users"
}
