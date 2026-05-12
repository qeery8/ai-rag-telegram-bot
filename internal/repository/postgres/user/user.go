package user

import (
	"github.com/jmoiron/sqlx"
)

type UserRepoSQL struct {
	db *sqlx.DB
}

func NewUserRepoSQL(db *sqlx.DB) *UserRepoSQL {
	return &UserRepoSQL{
		db: db,
	}
}
