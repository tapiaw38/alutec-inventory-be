package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucCategory "github.com/tapiaw38/alutec-inventory-be/internal/usecases/category"
)

func NewDeleteHandler(uc ucCategory.DeleteUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		out, appErr := uc.Execute(c, c.Param("id"))
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		message := "category deleted"
		if out.Archived {
			message = "category archived"
		}
		c.JSON(http.StatusOK, gin.H{"message": message, "data": out})
	}
}
