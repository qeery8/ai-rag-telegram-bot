package chat

import "github.com/jmoiron/sqlx"

type ChatRepoSQL struct {
	db *sqlx.DB
}

func NewChatRepoSQL(db *sqlx.DB) *ChatRepoSQL {
	return &ChatRepoSQL{
		db: db,
	}
}
