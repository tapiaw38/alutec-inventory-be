package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucCategory "github.com/tapiaw38/alutec-inventory-be/internal/usecases/category"
)

func NewGetHandler(uc ucCategory.GetUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		output, appErr := uc.Execute(c, c.Param("id"))
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
