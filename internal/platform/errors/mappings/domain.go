package mappings

import "net/http"

var (
	CategoryCreateError = ErrorDetails{
		InternalCode: "category:create:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to create category",
	}
	CategoryGetError = ErrorDetails{
		InternalCode: "category:get:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to get category",
	}
	CategoryNotFoundError = ErrorDetails{
		InternalCode: "category:get:not-found",
		StatusCode:   http.StatusNotFound,
		Message:      "category not found",
	}
	CategoryListError = ErrorDetails{
		InternalCode: "category:list:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to list categories",
	}
	CategoryUpdateError = ErrorDetails{
		InternalCode: "category:update:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to update category",
	}
	CategoryDeleteError = ErrorDetails{
		InternalCode: "category:delete:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to delete category",
	}

	SupplierCreateError = ErrorDetails{
		InternalCode: "supplier:create:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to create supplier",
	}
	SupplierGetError = ErrorDetails{
		InternalCode: "supplier:get:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to get supplier",
	}
	SupplierNotFoundError = ErrorDetails{
		InternalCode: "supplier:get:not-found",
		StatusCode:   http.StatusNotFound,
		Message:      "supplier not found",
	}
	SupplierListError = ErrorDetails{
		InternalCode: "supplier:list:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to list suppliers",
	}
	SupplierUpdateError = ErrorDetails{
		InternalCode: "supplier:update:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to update supplier",
	}
	SupplierDeleteError = ErrorDetails{
		InternalCode: "supplier:delete:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to delete supplier",
	}

	WarehouseCreateError = ErrorDetails{
		InternalCode: "warehouse:create:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to create warehouse",
	}
	WarehouseGetError = ErrorDetails{
		InternalCode: "warehouse:get:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to get warehouse",
	}
	WarehouseNotFoundError = ErrorDetails{
		InternalCode: "warehouse:get:not-found",
		StatusCode:   http.StatusNotFound,
		Message:      "warehouse not found",
	}
	WarehouseListError = ErrorDetails{
		InternalCode: "warehouse:list:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to list warehouses",
	}
	WarehouseUpdateError = ErrorDetails{
		InternalCode: "warehouse:update:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to update warehouse",
	}
	WarehouseDeleteError = ErrorDetails{
		InternalCode: "warehouse:delete:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to delete warehouse",
	}

	ProductCreateError = ErrorDetails{
		InternalCode: "product:create:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to create product",
	}
	ProductGetError = ErrorDetails{
		InternalCode: "product:get:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to get product",
	}
	ProductNotFoundError = ErrorDetails{
		InternalCode: "product:get:not-found",
		StatusCode:   http.StatusNotFound,
		Message:      "product not found",
	}
	ProductListError = ErrorDetails{
		InternalCode: "product:list:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to list products",
	}
	ProductUpdateError = ErrorDetails{
		InternalCode: "product:update:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to update product",
	}
	ProductDeleteError = ErrorDetails{
		InternalCode: "product:delete:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to delete product",
	}

	StockMovementCreateError = ErrorDetails{
		InternalCode: "stock_movement:create:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to create stock movement",
	}
	StockMovementGetError = ErrorDetails{
		InternalCode: "stock_movement:get:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to get stock movement",
	}
	StockMovementNotFoundError = ErrorDetails{
		InternalCode: "stock_movement:get:not-found",
		StatusCode:   http.StatusNotFound,
		Message:      "stock movement not found",
	}
	StockMovementListError = ErrorDetails{
		InternalCode: "stock_movement:list:error",
		StatusCode:   http.StatusInternalServerError,
		Message:      "failed to list stock movements",
	}
)
