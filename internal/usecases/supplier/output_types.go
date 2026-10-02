package supplier

import "github.com/tapiaw38/alutec-inventory-be/internal/domain"

type SupplierData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func toSupplierData(s domain.Supplier) SupplierData {
	return SupplierData{
		ID:        s.ID,
		Name:      s.Name,
		Phone:     s.Phone,
		Email:     s.Email,
		CreatedAt: s.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
