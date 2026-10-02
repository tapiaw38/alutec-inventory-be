package supplier

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Create(ctx context.Context, s domain.Supplier) (string, error) {
	var id string
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO suppliers (name, phone, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`, s.Name, s.Phone, s.Email).Scan(&id)
	return id, err
}
