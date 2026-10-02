package supplier

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucSupplier "github.com/tapiaw38/alutec-inventory-be/internal/usecases/supplier"
)

func NewCreateHandler(uc ucSupplier.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ucSupplier.CreateInput
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
