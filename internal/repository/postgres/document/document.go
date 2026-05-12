package document

import "github.com/jmoiron/sqlx"

type DocumentRepoSQL struct {
	db *sqlx.DB
}

func NewDocumentRepoSQL(db *sqlx.DB) *DocumentRepoSQL {
	return &DocumentRepoSQL{
		db: db,
	}
}
