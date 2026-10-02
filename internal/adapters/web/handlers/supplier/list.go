package supplier

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucSupplier "github.com/tapiaw38/alutec-inventory-be/internal/usecases/supplier"
)

func NewListHandler(uc ucSupplier.ListUsecase) gin.HandlerFunc {
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
