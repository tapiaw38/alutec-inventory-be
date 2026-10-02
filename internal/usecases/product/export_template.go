package product

import (
	"context"

	"github.com/xuri/excelize/v2"

	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
)

type (
	ExportTemplateUsecase interface {
		Execute(ctx context.Context) ([]byte, apperrors.ApplicationError)
	}

	exportTemplateUsecase struct{}
)

func NewExportTemplateUsecase() ExportTemplateUsecase {
	return &exportTemplateUsecase{}
}

var importTemplateHeaders = []string{
	"SKU", "Nombre", "Categoria", "Proveedor", "Unidad",
	"Precio Costo", "Precio Venta", "Stock Minimo", "Stock Actual",
}

func (u *exportTemplateUsecase) Execute(_ context.Context) ([]byte, apperrors.ApplicationError) {
	f := excelize.NewFile()
	defer f.Close()

	const sheet = "Productos"
	f.SetSheetName("Sheet1", sheet)

	for i, header := range importTemplateHeaders {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := f.SetCellValue(sheet, cell, header); err != nil {
			return nil, apperrors.NewInternalError(err)
		}
	}

	sample := []any{
		"PIN-2X4-3M", "Tabla de Pino 2x4 3m", "Maderas", "Maderera del Sur",
		"pieza", 4500, 6800, 20, 85,
	}
	for i, value := range sample {
		cell, _ := excelize.CoordinatesToCellName(i+1, 2)
		if err := f.SetCellValue(sheet, cell, value); err != nil {
			return nil, apperrors.NewInternalError(err)
		}
	}

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, apperrors.NewInternalError(err)
	}

	return buf.Bytes(), nil
}
