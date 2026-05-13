package document

import (
	"al-rag-tgbot/internal/domain"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *DocumentRepoSQL) FindDocumentByID(ctx context.Context, id uuid.UUID) (*domain.Document, error) {
	var document domain.Document
	const query = `
		SELECT * FROM documents WHERE id = $1
	`

	if err := r.db.GetContext(ctx, &document, query, id); err != nil {
		return nil, fmt.Errorf("documentRepo.FindByID: query failed: %w", err)
	}

	return &document, nil
}
