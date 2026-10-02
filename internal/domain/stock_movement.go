package domain

import "time"

type StockMovementType string

const (
	StockMovementTypeIn         StockMovementType = "in"
	StockMovementTypeOut        StockMovementType = "out"
	StockMovementTypeAdjustment StockMovementType = "adjustment"
)

// StockMovement is an append-only ledger entry: there is no update or delete
// repository operation for it by design. To correct a mistake, record a new
// movement that reverses it instead of mutating history.
type StockMovement struct {
	ID          string
	ProductID   string
	WarehouseID string
	Type        StockMovementType
	Quantity    int
	Date        time.Time
	Note        string
	CreatedAt   time.Time
}
