package user

import (
	"al-rag-tgbot/internal/domain"
	"context"
	"fmt"
)

func (r *UserRepoSQL) SaveUser(ctx context.Context, user *domain.User) error {
	const query = `
		INSERT INTO users (username)
		VALUES (:username)
		RETURNING id, created_at;
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("userRepo.SaveUser: prepare failed: %w", err)
	}
	defer stmt.Close()

	if err = stmt.GetContext(ctx, user, user); err != nil {
		return fmt.Errorf("userRepo.SaveUser: query failed: %w", err)
	}

	return nil
}
