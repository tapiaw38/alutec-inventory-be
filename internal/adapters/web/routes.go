package web

import (
	"github.com/gin-gonic/gin"
	handlerCategory "github.com/tapiaw38/alutec-inventory-be/internal/adapters/web/handlers/category"
	handlerProduct "github.com/tapiaw38/alutec-inventory-be/internal/adapters/web/handlers/product"
	handlerStockMovement "github.com/tapiaw38/alutec-inventory-be/internal/adapters/web/handlers/stock_movement"
	handlerSupplier "github.com/tapiaw38/alutec-inventory-be/internal/adapters/web/handlers/supplier"
	handlerWarehouse "github.com/tapiaw38/alutec-inventory-be/internal/adapters/web/handlers/warehouse"
	"github.com/tapiaw38/alutec-inventory-be/internal/usecases"
)

func RegisterRoutes(app *gin.Engine, uc *usecases.Usecases) {
	api := app.Group("/api")

	api.POST("/categories", handlerCategory.NewCreateHandler(uc.Category.Create))
	api.GET("/categories", handlerCategory.NewListHandler(uc.Category.List))
	api.GET("/categories/:id", handlerCategory.NewGetHandler(uc.Category.Get))
	api.PUT("/categories/:id", handlerCategory.NewUpdateHandler(uc.Category.Update))
	api.DELETE("/categories/:id", handlerCategory.NewDeleteHandler(uc.Category.Delete))

	api.POST("/suppliers", handlerSupplier.NewCreateHandler(uc.Supplier.Create))
	api.GET("/suppliers", handlerSupplier.NewListHandler(uc.Supplier.List))
	api.GET("/suppliers/:id", handlerSupplier.NewGetHandler(uc.Supplier.Get))
	api.PUT("/suppliers/:id", handlerSupplier.NewUpdateHandler(uc.Supplier.Update))
	api.DELETE("/suppliers/:id", handlerSupplier.NewDeleteHandler(uc.Supplier.Delete))

	api.POST("/warehouses", handlerWarehouse.NewCreateHandler(uc.Warehouse.Create))
	api.GET("/warehouses", handlerWarehouse.NewListHandler(uc.Warehouse.List))
	api.GET("/warehouses/:id", handlerWarehouse.NewGetHandler(uc.Warehouse.Get))
	api.PUT("/warehouses/:id", handlerWarehouse.NewUpdateHandler(uc.Warehouse.Update))
	api.DELETE("/warehouses/:id", handlerWarehouse.NewDeleteHandler(uc.Warehouse.Delete))

	api.POST("/products", handlerProduct.NewCreateHandler(uc.Product.Create))
	api.GET("/products", handlerProduct.NewListHandler(uc.Product.List))
	api.GET("/products/import-template", handlerProduct.NewExportTemplateHandler(uc.Product.ExportTemplate))
	api.POST("/products/import", handlerProduct.NewImportHandler(uc.Product.Import))
	api.GET("/products/:id", handlerProduct.NewGetHandler(uc.Product.Get))
	api.PUT("/products/:id", handlerProduct.NewUpdateHandler(uc.Product.Update))
	api.DELETE("/products/:id", handlerProduct.NewDeleteHandler(uc.Product.Delete))

	// Append-only ledger: intentionally no PUT or DELETE route here.
	api.POST("/stock-movements", handlerStockMovement.NewCreateHandler(uc.StockMovement.Create))
	api.GET("/stock-movements", handlerStockMovement.NewListHandler(uc.StockMovement.List))
}
