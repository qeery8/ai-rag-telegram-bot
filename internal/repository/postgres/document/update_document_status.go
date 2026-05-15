package document

import (
	"al-rag-tgbot/internal/domain"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *DocumentRepoSQL) UpdateDocumentStatus(ctx context.Context, id uuid.UUID, status domain.DocumentStatus) error {
	query := `
		UPDATE documents SET 
			status = :status
		WHERE id = :id 
	`

	args := map[string]any{
		"id":     id,
		"status": status,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("documentRepo.UpdateDocumentsStatus: prepare failed: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args)
	if err != nil {
		return fmt.Errorf("documentRepo.UpdateDocumentStatus: query failed: %w", err)
	}

	return nil
}
