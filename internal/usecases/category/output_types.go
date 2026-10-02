package category

import "github.com/tapiaw38/alutec-inventory-be/internal/domain"

type CategoryData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}

func toCategoryData(c domain.Category) CategoryData {
	return CategoryData{
		ID:        c.ID,
		Name:      c.Name,
		Type:      string(c.Type),
		CreatedAt: c.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
