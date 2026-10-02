package stockmovement

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucStockMovement "github.com/tapiaw38/alutec-inventory-be/internal/usecases/stock_movement"
)

func NewListHandler(uc ucStockMovement.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := ucStockMovement.ListFilter{
			ProductID:   c.Query("product_id"),
			CategoryID:  c.Query("category_id"),
			WarehouseID: c.Query("warehouse_id"),
			Type:        c.Query("type"),
			Search:      c.Query("search"),
			DateFrom:    c.Query("date_from"),
			DateTo:      c.Query("date_to"),
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
