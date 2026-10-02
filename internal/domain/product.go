package domain

import "time"

type Product struct {
	ID         string
	SKU        string
	Name       string
	CategoryID string
	SupplierID string
	Unit       string
	CostPrice  float64
	SalePrice  float64
	MinStock   int
	StockQty   int
	ImageURL   string
	CreatedAt  time.Time
}
