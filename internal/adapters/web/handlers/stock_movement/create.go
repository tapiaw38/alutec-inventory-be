package stockmovement

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucStockMovement "github.com/tapiaw38/alutec-inventory-be/internal/usecases/stock_movement"
)

func NewCreateHandler(uc ucStockMovement.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ucStockMovement.CreateInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "common:bad-request", "message": err.Error()})
			return
		}

		output, appErr := uc.Execute(c, input)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusCreated, output)
	}
}
