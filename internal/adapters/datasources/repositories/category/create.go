package category

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Create(ctx context.Context, c domain.Category) (string, error) {
	var id string
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO categories (name, type)
		VALUES ($1, $2)
		RETURNING id
	`, c.Name, c.Type).Scan(&id)
	return id, err
}
