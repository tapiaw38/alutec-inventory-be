package usecases

import (
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	ucAuth "github.com/tapiaw38/alutec-inventory-be/internal/usecases/auth"
	ucCategory "github.com/tapiaw38/alutec-inventory-be/internal/usecases/category"
	ucProduct "github.com/tapiaw38/alutec-inventory-be/internal/usecases/product"
	ucStockMovement "github.com/tapiaw38/alutec-inventory-be/internal/usecases/stock_movement"
	ucSupplier "github.com/tapiaw38/alutec-inventory-be/internal/usecases/supplier"
	ucWarehouse "github.com/tapiaw38/alutec-inventory-be/internal/usecases/warehouse"
)

type AuthUsecases struct {
	Login ucAuth.LoginUsecase
}

type CategoryUsecases struct {
	Create ucCategory.CreateUsecase
	Get    ucCategory.GetUsecase
	List   ucCategory.ListUsecase
	Update ucCategory.UpdateUsecase
	Delete ucCategory.DeleteUsecase
}

type SupplierUsecases struct {
	Create ucSupplier.CreateUsecase
	Get    ucSupplier.GetUsecase
	List   ucSupplier.ListUsecase
	Update ucSupplier.UpdateUsecase
	Delete ucSupplier.DeleteUsecase
}

type WarehouseUsecases struct {
	Create ucWarehouse.CreateUsecase
	Get    ucWarehouse.GetUsecase
	List   ucWarehouse.ListUsecase
	Update ucWarehouse.UpdateUsecase
	Delete ucWarehouse.DeleteUsecase
}

type ProductUsecases struct {
	Create         ucProduct.CreateUsecase
	Get            ucProduct.GetUsecase
	List           ucProduct.ListUsecase
	Update         ucProduct.UpdateUsecase
	Delete         ucProduct.DeleteUsecase
	Import         ucProduct.ImportUsecase
	ExportTemplate ucProduct.ExportTemplateUsecase
}

type StockMovementUsecases struct {
	Create ucStockMovement.CreateUsecase
	List   ucStockMovement.ListUsecase
}

type Usecases struct {
	Auth          AuthUsecases
	Category      CategoryUsecases
	Supplier      SupplierUsecases
	Warehouse     WarehouseUsecases
	Product       ProductUsecases
	StockMovement StockMovementUsecases
}

func NewUsecases(contextFactory appcontext.Factory) *Usecases {
	return &Usecases{
		Auth: AuthUsecases{
			Login: ucAuth.NewLoginUsecase(),
		},
		Category: CategoryUsecases{
			Create: ucCategory.NewCreateUsecase(contextFactory),
			Get:    ucCategory.NewGetUsecase(contextFactory),
			List:   ucCategory.NewListUsecase(contextFactory),
			Update: ucCategory.NewUpdateUsecase(contextFactory),
			Delete: ucCategory.NewDeleteUsecase(contextFactory),
		},
		Supplier: SupplierUsecases{
			Create: ucSupplier.NewCreateUsecase(contextFactory),
			Get:    ucSupplier.NewGetUsecase(contextFactory),
			List:   ucSupplier.NewListUsecase(contextFactory),
			Update: ucSupplier.NewUpdateUsecase(contextFactory),
			Delete: ucSupplier.NewDeleteUsecase(contextFactory),
		},
		Warehouse: WarehouseUsecases{
			Create: ucWarehouse.NewCreateUsecase(contextFactory),
			Get:    ucWarehouse.NewGetUsecase(contextFactory),
			List:   ucWarehouse.NewListUsecase(contextFactory),
			Update: ucWarehouse.NewUpdateUsecase(contextFactory),
			Delete: ucWarehouse.NewDeleteUsecase(contextFactory),
		},
		Product: func() ProductUsecases {
			createProduct := ucProduct.NewCreateUsecase(contextFactory)
			return ProductUsecases{
				Create:         createProduct,
				Get:            ucProduct.NewGetUsecase(contextFactory),
				List:           ucProduct.NewListUsecase(contextFactory),
				Update:         ucProduct.NewUpdateUsecase(contextFactory),
				Delete:         ucProduct.NewDeleteUsecase(contextFactory),
				Import:         ucProduct.NewImportUsecase(contextFactory, createProduct),
				ExportTemplate: ucProduct.NewExportTemplateUsecase(),
			}
		}(),
		StockMovement: StockMovementUsecases{
			Create: ucStockMovement.NewCreateUsecase(contextFactory),
			List:   ucStockMovement.NewListUsecase(contextFactory),
		},
	}
}
