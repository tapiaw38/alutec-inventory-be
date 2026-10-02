package warehouse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucWarehouse "github.com/tapiaw38/alutec-inventory-be/internal/usecases/warehouse"
)

func NewListHandler(uc ucWarehouse.ListUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		output, appErr := uc.Execute(c)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
