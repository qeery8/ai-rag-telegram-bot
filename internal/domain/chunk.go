package domain

import (
	"time"

	"github.com/google/uuid"
)

type Chunk struct {
	ID         uuid.UUID `db:"id"`
	DocumentID uuid.UUID `db:"document_id"`
	Content    string    `db:"content"`
	Position   int       `db:"position"`
	VectorID   string    `db:"vector_id"`
	CreatedAt  time.Time `db:"created_at"`
}
