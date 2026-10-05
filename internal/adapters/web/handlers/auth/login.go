package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	ucAuth "github.com/tapiaw38/alutec-inventory-be/internal/usecases/auth"
)

func NewLoginHandler(uc ucAuth.LoginUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ucAuth.LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			appErr := apperrors.NewBadRequestError("email y contraseña son obligatorios")
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		out, appErr := uc.Execute(c, input)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, out)
	}
}
