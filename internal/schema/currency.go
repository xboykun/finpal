package schema

import "time"

type Currency struct {
	ID        int64
	Name      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Unit struct {
	ID        int64
	Name      string
	Value     string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LinkedCurrency struct {
	ID         int64
	CurrencyID int64
	UnitID     int64
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
