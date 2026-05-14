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
	ID        uuid.UUID `db:"id"`
	UserID    int64     `db:"user_id"`
	Role      ChatRole  `db:"role"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
