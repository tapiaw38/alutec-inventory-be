package category

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Update(ctx context.Context, id string, c domain.Category) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE categories SET name = $1, type = $2 WHERE id = $3
	`, c.Name, c.Type, id)
	return err
}
