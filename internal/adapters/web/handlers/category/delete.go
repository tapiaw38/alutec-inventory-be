package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucCategory "github.com/tapiaw38/alutec-inventory-be/internal/usecases/category"
)

func NewDeleteHandler(uc ucCategory.DeleteUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if appErr := uc.Execute(c, c.Param("id")); appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
	}
}
