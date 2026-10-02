package domain

import "time"

type CategoryType string

const (
	CategoryTypeRawMaterial  CategoryType = "raw_material"
	CategoryTypeFinishedGood CategoryType = "finished_good"
	CategoryTypeTool         CategoryType = "tool"
)

type Category struct {
	ID        string
	Name      string
	Type      CategoryType
	CreatedAt time.Time
}
