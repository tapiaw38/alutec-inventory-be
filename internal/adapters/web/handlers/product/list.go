package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucProduct "github.com/tapiaw38/alutec-inventory-be/internal/usecases/product"
)

func NewListHandler(uc ucProduct.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := ucProduct.ListFilter{
			CategoryID: c.Query("category_id"),
			SupplierID: c.Query("supplier_id"),
			Search:     c.Query("search"),
			LowStock:   c.Query("low_stock") == "true",
		}

		output, appErr := uc.Execute(c, filter)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
