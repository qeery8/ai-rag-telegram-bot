package domain

import (
	"time"

	"github.com/google/uuid"
)

type ChatRole string

const (
	RoleUser  ChatRole = "user"
	RoleAdmin ChatRole = "admin"
)

type Chat struct {
	ID        uuid.UUID
	UserID    int64
	Role      ChatRole
	Content   string
	CreatedAt time.Time
}
