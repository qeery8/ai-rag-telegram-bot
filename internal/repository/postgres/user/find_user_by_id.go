package user

import (
	"al-rag-tgbot/internal/domain"
	"context"
	"fmt"
)

func (r *UserRepoSQL) FindByID(ctx context.Context, userID int64) (*domain.User, error) {
	var user domain.User
	const query = `
		SELECT * FROM users WHERE id = $1
	`

	if err := r.db.GetContext(ctx, &user, query, userID); err != nil {
		return nil, fmt.Errorf("userRepo.FindByID: query failed: %w", err)
	}

	return &user, nil
}
