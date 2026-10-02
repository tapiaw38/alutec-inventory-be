package domain

import "time"

type Supplier struct {
	ID        string
	Name      string
	Phone     string
	Email     string
	CreatedAt time.Time
}
