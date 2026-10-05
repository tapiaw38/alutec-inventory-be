package supplier

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucSupplier "github.com/tapiaw38/alutec-inventory-be/internal/usecases/supplier"
)

func NewDeleteHandler(uc ucSupplier.DeleteUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		out, appErr := uc.Execute(c, c.Param("id"))
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		message := "supplier deleted"
		if out.Archived {
			message = "supplier archived"
		}
		c.JSON(http.StatusOK, gin.H{"message": message, "data": out})
	}
}
