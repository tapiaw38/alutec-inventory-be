package supplier

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) List(ctx context.Context) ([]domain.Supplier, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, COALESCE(phone, ''), COALESCE(email, ''), created_at
		FROM suppliers
		WHERE archived_at IS NULL
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	suppliers := []domain.Supplier{}
	for rows.Next() {
		var s domain.Supplier
		if err := rows.Scan(&s.ID, &s.Name, &s.Phone, &s.Email, &s.CreatedAt); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, s)
	}
	return suppliers, rows.Err()
}
