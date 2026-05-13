package document

import (
	"al-rag-tgbot/internal/domain"
	"context"
	"fmt"
)

func (r *DocumentRepoSQL) SaveDocument(ctx context.Context, document *domain.Document) error {
	const query = `
		INSERT INTO documents (user_id, name, type, status)
		VALUES (:user_id, :name, :type, :status)
		RETURNING id, created_at;
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("documentRepo.SaveDocument: prepare failed: %w", err)
	}
	stmt.Close()

	if err = stmt.GetContext(ctx, document, document); err != nil {
		return fmt.Errorf("documentRepo.SaveDocument: query failed: %w", err)
	}

	return nil
}
