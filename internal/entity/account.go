package entity

import (
	"time"

	"github.com/google/uuid"
)

type AccountSchema struct {
	ID        uuid.UUID
	PartnerID uuid.UUID
	Name      string
	IsStatus  int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateAccountInput struct {
	Payload *AccountBodyJson `in:"body=json"`
}

type AccountBodyJson struct {
	PartnerCode string `json:"partnerCode"`
	Name        string `json:"name"`
}
