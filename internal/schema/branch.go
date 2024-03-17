package schema

import "time"

type Branch struct {
	ID             int64
	OrganizationID int64
	Name           string
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
