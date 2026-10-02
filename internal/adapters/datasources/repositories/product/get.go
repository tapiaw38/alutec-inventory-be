package product

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Get(ctx context.Context, id string) (*domain.Product, error) {
	var p domain.Product
	err := r.db.QueryRowContext(ctx, `
		SELECT id, sku, name, category_id, supplier_id, unit, cost_price, sale_price, min_stock, stock_qty, COALESCE(image_url, ''), created_at
		FROM products
		WHERE id = $1
	`, id).Scan(&p.ID, &p.SKU, &p.Name, &p.CategoryID, &p.SupplierID, &p.Unit, &p.CostPrice, &p.SalePrice, &p.MinStock, &p.StockQty, &p.ImageURL, &p.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}
