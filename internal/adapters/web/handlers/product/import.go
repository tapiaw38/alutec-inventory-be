package product

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	ucProduct "github.com/tapiaw38/alutec-inventory-be/internal/usecases/product"
)

func NewImportHandler(uc ucProduct.ImportUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "common:bad-request", "message": "falta el archivo"})
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "common:bad-request", "message": "no se pudo abrir el archivo"})
			return
		}
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "common:bad-request", "message": "no se pudo leer el archivo"})
			return
		}

		output, appErr := uc.Execute(c, fileBytes, fileHeader.Filename)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
