package product

import (
	"context"
	"fmt"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) List(ctx context.Context, filter ListFilterOptions) ([]domain.Product, error) {
	query := `
		SELECT id, sku, name, category_id, supplier_id, unit, cost_price, sale_price, min_stock, stock_qty, COALESCE(image_url, ''), created_at
		FROM products
		WHERE archived_at IS NULL
	`
	args := []any{}

	if filter.CategoryID != "" {
		args = append(args, filter.CategoryID)
		query += fmt.Sprintf(" AND category_id = $%d", len(args))
	}
	if filter.SupplierID != "" {
		args = append(args, filter.SupplierID)
		query += fmt.Sprintf(" AND supplier_id = $%d", len(args))
	}
	if filter.Search != "" {
		args = append(args, "%"+filter.Search+"%")
		query += fmt.Sprintf(" AND (name ILIKE $%d OR sku ILIKE $%d)", len(args), len(args))
	}
	if filter.LowStock {
		query += " AND stock_qty <= min_stock"
	}
	query += " ORDER BY name ASC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.SKU, &p.Name, &p.CategoryID, &p.SupplierID, &p.Unit, &p.CostPrice, &p.SalePrice, &p.MinStock, &p.StockQty, &p.ImageURL, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}
