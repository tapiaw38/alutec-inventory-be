package warehouse

import "github.com/tapiaw38/alutec-inventory-be/internal/domain"

type WarehouseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	CreatedAt string `json:"created_at"`
}

func toWarehouseData(w domain.Warehouse) WarehouseData {
	return WarehouseData{
		ID:        w.ID,
		Name:      w.Name,
		Location:  w.Location,
		CreatedAt: w.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
