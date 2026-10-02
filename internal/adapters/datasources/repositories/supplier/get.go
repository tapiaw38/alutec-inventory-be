package supplier

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Get(ctx context.Context, id string) (*domain.Supplier, error) {
	var s domain.Supplier
	err := r.db.QueryRowContext(ctx, `
		SELECT id, name, COALESCE(phone, ''), COALESCE(email, ''), created_at
		FROM suppliers
		WHERE id = $1
	`, id).Scan(&s.ID, &s.Name, &s.Phone, &s.Email, &s.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}
