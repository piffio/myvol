package models

import (
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
)

type User struct {
	ID         uuid.UUID    `json:"id" db:"id"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	Name       string       `json:"name" db:"name"`
	Email      nulls.String `json:"email" db:"email"`
	Provider   string       `json:"provider" db:"provider"`
	ProviderID string       `json:"provider_id" db:"provider_id"`
}
