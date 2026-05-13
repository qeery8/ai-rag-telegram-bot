package document

import (
	"al-rag-tgbot/internal/domain"
	"context"
	"fmt"
)

func (r *DocumentRepoSQL) FindDocumentByUserID(ctx context.Context, userID int64) (*[]domain.Document, error) {
	var documents []domain.Document
	const query = `
		SELECT * FROM documents WHERE user_id = $1
	`

	if err := r.db.GetContext(ctx, &documents, query, userID); err != nil {
		return nil, fmt.Errorf("documentRepo.FindDocumentByID: query failed: %w", err)
	}

	return &documents, nil
}
