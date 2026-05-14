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
	ID        uuid.UUID      `db:"id"`
	UserID    int64          `db:"user_id"`
	Name      string         `db:"name"`
	Type      string         `db:"type"`
	Status    DocumentStatus `db:"status"`
	CreatedAt time.Time      `db:"created_at"`
}
