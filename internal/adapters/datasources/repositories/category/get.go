package category

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Get(ctx context.Context, id string) (*domain.Category, error) {
	var c domain.Category
	err := r.db.QueryRowContext(ctx, `
		SELECT id, name, type, created_at
		FROM categories
		WHERE id = $1
	`, id).Scan(&c.ID, &c.Name, &c.Type, &c.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}
