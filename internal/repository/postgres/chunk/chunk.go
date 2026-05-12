package chunk

import "github.com/jmoiron/sqlx"

type ChunkRepoSQL struct {
	db *sqlx.DB
}

func NewChunkRepoSQL(db *sqlx.DB) *ChunkRepoSQL {
	return &ChunkRepoSQL{
		db: db,
	}
}
