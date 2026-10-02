package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucProduct "github.com/tapiaw38/alutec-inventory-be/internal/usecases/product"
)

func NewCreateHandler(uc ucProduct.CreateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ucProduct.CreateInput
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
