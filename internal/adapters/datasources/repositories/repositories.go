package repositories

import (
	"github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources"
	"github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources/repositories/category"
	"github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources/repositories/product"
	stockmovement "github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources/repositories/stock_movement"
	"github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources/repositories/supplier"
	"github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources/repositories/warehouse"
)

type Repositories struct {
	Category      category.Repository
	Supplier      supplier.Repository
	Warehouse     warehouse.Repository
	Product       product.Repository
	StockMovement stockmovement.Repository
}

type Factory func() *Repositories

func NewFactory(ds *datasources.Datasources) Factory {
	return func() *Repositories {
		return &Repositories{
			Category:      category.NewRepository(ds.DB),
			Supplier:      supplier.NewRepository(ds.DB),
			Warehouse:     warehouse.NewRepository(ds.DB),
			Product:       product.NewRepository(ds.DB),
			StockMovement: stockmovement.NewRepository(ds.DB),
		}
	}
}
