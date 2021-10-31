package model

import (
	"database/sql"
	"time"
)

type User struct {
	ChatID        int64         `db:"chat_id"`
	Available     bool          `db:"available"`
	LastActivity  time.Time     `db:"last_activity"`
	MatchChatID   sql.NullInt64 `db:"match_chat_id"`
	RegisterDate  time.Time     `db:"register_date"`
	PreviousMatch sql.NullInt64 `db:"previous_match"`
	AllowPictures bool          `db:"allow_pictures"`
	BannedUntil   sql.NullTime  `db:"banned_until"`
}

func (p *User) TableName() string {
	return "users"
}
