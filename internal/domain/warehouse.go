package domain

import "time"

type Warehouse struct {
	ID        string
	Name      string
	Location  string
	CreatedAt time.Time
}
