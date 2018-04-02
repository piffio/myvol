package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

type Device struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Model       string    `json:"model" db:"model"`
	Description string    `json:"description" db:"description"`
	Serial      string    `json:"serial" db:"serial"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
}

// Devices is not required by pop and may be deleted
type Devices []Device
