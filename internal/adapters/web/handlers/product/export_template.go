package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ucProduct "github.com/tapiaw38/alutec-inventory-be/internal/usecases/product"
)

func NewExportTemplateHandler(uc ucProduct.ExportTemplateUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileBytes, appErr := uc.Execute(c)
		if appErr != nil {
			appErr.Log(c)
			c.JSON(appErr.StatusCode(), appErr)
			return
		}

		c.Header("Content-Disposition", `attachment; filename="plantilla-productos.xlsx"`)
		c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
	}
}
