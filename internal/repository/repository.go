package repository

import (
	"al-rag-tgbot/internal/domain"
	"context"

	"github.com/google/uuid"
)

type Repo struct {
	UserRepository
	DocumentRepository
	ChatRepository
	ChunkRepository
}

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, userID int64) (*domain.User, error)
}

type DocumentRepository interface {
	Save(ctx context.Context, document *domain.Document) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Document, error)
	FindByUserID(ctx context.Context, userID int64) ([]*domain.Document, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status domain.DocumentStatus) error
}

type ChatRepository interface {
	Save(ctx context.Context, msg *domain.Chat) error
	FindByUserID(ctx context.Context, userID int64, limit int) ([]*domain.Chat, error)
}

type ChunkRepository interface {
	Save(ctx context.Context, chunk *domain.Chunk) error
	SaveBatch(ctx context.Context, chunks []*domain.Chunk) error
	FindByDocumentID(ctx context.Context, document uuid.UUID) ([]*domain.Chunk, error)
	UpdateVector(ctx context.Context, id uuid.UUID, vectorID string) error
}

func NewRepo(
	userRepo UserRepository,
	documentRepo DocumentRepository,
	chatRepo ChatRepository,
	chunkRepo ChunkRepository,
) *Repo {
	return &Repo{
		UserRepository:     userRepo,
		DocumentRepository: documentRepo,
		ChatRepository:     chatRepo,
		ChunkRepository:    chunkRepo,
	}
}
