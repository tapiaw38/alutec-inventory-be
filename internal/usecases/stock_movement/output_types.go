package stockmovement

import "github.com/tapiaw38/alutec-inventory-be/internal/domain"

type StockMovementData struct {
	ID          string `json:"id"`
	ProductID   string `json:"product_id"`
	WarehouseID string `json:"warehouse_id"`
	Type        string `json:"type"`
	Quantity    int    `json:"quantity"`
	Date        string `json:"date"`
	Note        string `json:"note"`
	CreatedAt   string `json:"created_at"`
}

func toStockMovementData(m domain.StockMovement) StockMovementData {
	return StockMovementData{
		ID:          m.ID,
		ProductID:   m.ProductID,
		WarehouseID: m.WarehouseID,
		Type:        string(m.Type),
		Quantity:    m.Quantity,
		Date:        m.Date.Format("2006-01-02"),
		Note:        m.Note,
		CreatedAt:   m.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
