package supplier

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Update(ctx context.Context, id string, s domain.Supplier) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE suppliers SET name = $1, phone = $2, email = $3 WHERE id = $4
	`, s.Name, s.Phone, s.Email, id)
	return err
}
