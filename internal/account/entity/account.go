package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	PartnerID int64
	Name      string
	IsStatus  int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountJsonBody struct {
	PartnerID int64  `json:"partnerId"`
	Name      string `json:"name"`
}

type AccountInput struct {
	Body *AccountJsonBody `in:"body=json"`
}

func X()  {
	uuid.NewV7()
}