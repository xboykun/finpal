package schema

import "time"

type Transaction struct {
	ID        int64
	RefID     string
	Detail    []Goods
	Total     float64
	Status    string
	CreatedAt time.Time
}

type Goods struct {
	Name     string
	Quantity int32
	Amount   float64
}

type History struct {
	ID             int64
	TransactionID  int64
	OrganizationID int64
	BranchID       int64
	CustomerID     int64
	CurrencyID     int64
	Total          float64
	PrevBalance    float64
	CurrBalance    float64
	Status         string
	CreatedAt      time.Time
}
