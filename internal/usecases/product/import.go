package product

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/xuri/excelize/v2"

	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
)

type (
	ImportUsecase interface {
		Execute(ctx context.Context, fileBytes []byte, filename string) (*ImportOutput, apperrors.ApplicationError)
	}

	importUsecase struct {
		contextFactory appcontext.Factory
		createUsecase  CreateUsecase
	}

	ImportRowError struct {
		Row     int    `json:"row"`
		Message string `json:"message"`
	}

	ImportOutput struct {
		Imported int              `json:"imported"`
		Errors   []ImportRowError `json:"errors"`
	}
)

func NewImportUsecase(contextFactory appcontext.Factory, createUsecase CreateUsecase) ImportUsecase {
	return &importUsecase{contextFactory: contextFactory, createUsecase: createUsecase}
}

// importHeaderKey maps a normalized spreadsheet header to the internal field
// it feeds. Normalization strips accents, spaces and casing.
var importHeaderKey = map[string]string{
	"sku":           "sku",
	"nombre":        "name",
	"categoria":     "category",
	"proveedor":     "supplier",
	"unidad":        "unit",
	"preciocosto":   "cost_price",
	"preciodecosto": "cost_price",
	"precioventa":   "sale_price",
	"preciodeventa": "sale_price",
	"stockminimo":   "min_stock",
	"stockactual":   "stock_qty",
}

func (u *importUsecase) Execute(ctx context.Context, fileBytes []byte, filename string) (*ImportOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	rawRows, err := readSpreadsheetRows(fileBytes, filename)
	if err != nil {
		return nil, apperrors.NewBadRequestError("no se pudo leer el archivo: " + err.Error())
	}
	if len(rawRows) < 2 {
		return &ImportOutput{Imported: 0, Errors: nil}, nil
	}

	header := rawRows[0]
	fieldByColumn := make(map[int]string, len(header))
	for col, raw := range header {
		if key, ok := importHeaderKey[normalizeHeader(raw)]; ok {
			fieldByColumn[col] = key
		}
	}

	categories, err := app.Repositories.Category.List(ctx)
	if err != nil {
		return nil, apperrors.NewInternalError(err)
	}
	suppliers, err := app.Repositories.Supplier.List(ctx)
	if err != nil {
		return nil, apperrors.NewInternalError(err)
	}

	categoryIDByName := make(map[string]string, len(categories))
	for _, c := range categories {
		categoryIDByName[normalizeText(c.Name)] = c.ID
	}
	supplierIDByName := make(map[string]string, len(suppliers))
	for _, s := range suppliers {
		supplierIDByName[normalizeText(s.Name)] = s.ID
	}

	output := &ImportOutput{}

	for i, raw := range rawRows[1:] {
		rowNumber := i + 2 // header is row 1, spreadsheets are 1-indexed

		entry := make(map[string]string, len(fieldByColumn))
		for col, field := range fieldByColumn {
			if col < len(raw) {
				entry[field] = strings.TrimSpace(raw[col])
			}
		}

		if entry["sku"] == "" && entry["name"] == "" {
			continue // skip fully blank trailing rows
		}
		if entry["sku"] == "" || entry["name"] == "" {
			output.Errors = append(output.Errors, ImportRowError{Row: rowNumber, Message: "Falta SKU o Nombre"})
			continue
		}

		categoryID, ok := categoryIDByName[normalizeText(entry["category"])]
		if !ok {
			output.Errors = append(output.Errors, ImportRowError{
				Row: rowNumber, Message: fmt.Sprintf("Categoría no encontrada: %q", entry["category"]),
			})
			continue
		}
		supplierID, ok := supplierIDByName[normalizeText(entry["supplier"])]
		if !ok {
			output.Errors = append(output.Errors, ImportRowError{
				Row: rowNumber, Message: fmt.Sprintf("Proveedor no encontrado: %q", entry["supplier"]),
			})
			continue
		}

		unit := entry["unit"]
		if unit == "" {
			unit = "pieza"
		}

		_, appErr := u.createUsecase.Execute(ctx, CreateInput{
			SKU:        entry["sku"],
			Name:       entry["name"],
			CategoryID: categoryID,
			SupplierID: supplierID,
			Unit:       unit,
			CostPrice:  parseFloat(entry["cost_price"]),
			SalePrice:  parseFloat(entry["sale_price"]),
			MinStock:   parseInt(entry["min_stock"]),
			StockQty:   parseInt(entry["stock_qty"]),
		})
		if appErr != nil {
			output.Errors = append(output.Errors, ImportRowError{Row: rowNumber, Message: appErr.Message()})
			continue
		}

		output.Imported++
	}

	return output, nil
}

func readSpreadsheetRows(fileBytes []byte, filename string) ([][]string, error) {
	if strings.HasSuffix(strings.ToLower(filename), ".csv") {
		return readCSVRows(fileBytes)
	}

	f, err := excelize.OpenReader(bytes.NewReader(fileBytes))
	if err != nil {
		return readCSVRows(fileBytes) // fall back for mislabeled CSVs
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("the workbook has no sheets")
	}
	return f.GetRows(sheets[0])
}

func readCSVRows(fileBytes []byte) ([][]string, error) {
	reader := csv.NewReader(bytes.NewReader(fileBytes))
	reader.FieldsPerRecord = -1
	return reader.ReadAll()
}

func normalizeHeader(value string) string {
	var b strings.Builder
	for _, r := range normalizeText(value) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func normalizeText(value string) string {
	return strings.ToLower(strings.TrimSpace(stripDiacritics(value)))
}

func stripDiacritics(value string) string {
	replacer := strings.NewReplacer(
		"á", "a", "é", "e", "í", "i", "ó", "o", "ú", "u", "ñ", "n",
		"Á", "a", "É", "e", "Í", "i", "Ó", "o", "Ú", "u", "Ñ", "n",
	)
	return replacer.Replace(value)
}

func parseFloat(value string) float64 {
	normalized := strings.ReplaceAll(value, ",", ".")
	f, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return 0
	}
	return f
}

func parseInt(value string) int {
	i, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		f := parseFloat(value)
		return int(f)
	}
	return i
}
