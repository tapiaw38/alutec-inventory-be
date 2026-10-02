package product

import "github.com/tapiaw38/alutec-inventory-be/internal/domain"

type ProductData struct {
	ID         string  `json:"id"`
	SKU        string  `json:"sku"`
	Name       string  `json:"name"`
	CategoryID string  `json:"category_id"`
	SupplierID string  `json:"supplier_id"`
	Unit       string  `json:"unit"`
	CostPrice  float64 `json:"cost_price"`
	SalePrice  float64 `json:"sale_price"`
	MinStock   int     `json:"min_stock"`
	StockQty   int     `json:"stock_qty"`
	ImageURL   string  `json:"image_url"`
	CreatedAt  string  `json:"created_at"`
}

func toProductData(p domain.Product) ProductData {
	return ProductData{
		ID:         p.ID,
		SKU:        p.SKU,
		Name:       p.Name,
		CategoryID: p.CategoryID,
		SupplierID: p.SupplierID,
		Unit:       p.Unit,
		CostPrice:  p.CostPrice,
		SalePrice:  p.SalePrice,
		MinStock:   p.MinStock,
		StockQty:   p.StockQty,
		ImageURL:   p.ImageURL,
		CreatedAt:  p.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
