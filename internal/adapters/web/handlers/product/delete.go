package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucProduct "github.com/tapiaw38/alutec-inventory-be/internal/usecases/product"
)

func NewDeleteHandler(uc ucProduct.DeleteUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		out, appErr := uc.Execute(c, c.Param("id"))
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		message := "product deleted"
		if out.Archived {
			message = "product archived"
		}
		c.JSON(http.StatusOK, gin.H{"message": message, "data": out})
	}
}
