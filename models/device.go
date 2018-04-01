package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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

// String is not required by pop and may be deleted
func (d Device) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Devices is not required by pop and may be deleted
type Devices []Device

// String is not required by pop and may be deleted
func (d Devices) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Device) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: d.Model, Name: "Model"},
		&validators.StringIsPresent{Field: d.Description, Name: "Description"},
		&validators.StringIsPresent{Field: d.Serial, Name: "Serial"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Device) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Device) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
