package postgres

import (
	"al-rag-tgbot/internal/config"
	"al-rag-tgbot/internal/repository/postgres/chat"
	"al-rag-tgbot/internal/repository/postgres/chunk"
	"al-rag-tgbot/internal/repository/postgres/document"
	"al-rag-tgbot/internal/repository/postgres/user"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	*user.UserRepoSQL
	*document.DocumentRepoSQL
	*chunk.ChunkRepoSQL
	*chat.ChatRepoSQL
}

func NewRepo(db *sqlx.DB) *Postgres {
	return &Postgres{
		UserRepoSQL:     user.NewUserRepoSQL(db),
		DocumentRepoSQL: document.NewDocumentRepoSQL(db),
		ChunkRepoSQL:    chunk.NewChunkRepoSQL(db),
		ChatRepoSQL:     chat.NewChatRepoSQL(db),
	}
}

func NewPostgresConnection(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.DbName, cfg.Postgres.Password, cfg.Postgres.SslMode,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("connection could not be established: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("couldn't connect to db: %w", err)
	}

	return db, nil
}
