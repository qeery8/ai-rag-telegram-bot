package domain

import (
	"time"

	"github.com/google/uuid"
)

type DocumentStatus string

const (
	StatusPending    DocumentStatus = "pending"
	StatusProcessing DocumentStatus = "processing"
	StatusDone       DocumentStatus = "done"
	StatusFailed     DocumentStatus = "failed"
)

type Document struct {
	ID        uuid.UUID
	UserID    int64
	Name      string
	Type      string
	Status    DocumentStatus
	CreatedAt time.Time
}
