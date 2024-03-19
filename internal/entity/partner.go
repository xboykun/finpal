package entity

import "time"

type PartnerSchema struct {
	ID        int64
	Name      string
	IsStatus  int8
	CreatedAt time.Time
	UpdatedAt time.Time
}
