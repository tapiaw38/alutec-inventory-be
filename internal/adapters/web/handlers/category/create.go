package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucCategory "github.com/tapiaw38/alutec-inventory-be/internal/usecases/category"
)

func NewCreateHandler(uc ucCategory.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ucCategory.CreateInput
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
