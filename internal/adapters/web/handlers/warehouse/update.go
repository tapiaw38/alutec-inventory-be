package warehouse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucWarehouse "github.com/tapiaw38/alutec-inventory-be/internal/usecases/warehouse"
)

func NewUpdateHandler(uc ucWarehouse.UpdateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ucWarehouse.UpdateInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "common:bad-request", "message": err.Error()})
			return
		}

		output, appErr := uc.Execute(c, c.Param("id"), input)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
