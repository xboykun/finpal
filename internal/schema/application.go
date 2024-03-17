package schema

import "time"

type Application struct {
	ID             int64
	OrganizationID int64
	BranchID       int64
	Name           string
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
