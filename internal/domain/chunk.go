package domain

import (
	"time"

	"github.com/google/uuid"
)

type Chunk struct {
	ID         uuid.UUID
	DocumentID uuid.UUID
	Content    string
	Position   int
	VectorID   string
	CreatedAt  time.Time
}
