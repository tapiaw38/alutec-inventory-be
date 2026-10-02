package warehouse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucWarehouse "github.com/tapiaw38/alutec-inventory-be/internal/usecases/warehouse"
)

func NewDeleteHandler(uc ucWarehouse.DeleteUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if appErr := uc.Execute(c, c.Param("id")); appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "warehouse deleted"})
	}
}
