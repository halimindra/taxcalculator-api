package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/halimindra/taxcalculator-api/services"
)

type Tax struct {
	ID        int                       `json:"id" db:"id"`
	CreatedAt time.Time                 `json:"-" db:"created_at"`
	UpdatedAt time.Time                 `json:"-" db:"updated_at"`
	Name      string                    `json:"name" db:"name" validate:"required"`
	TaxCode   string                    `json:"tax_code" db:"tax_code" validate:"required,oneof=1 2 3"`
	Price     float64                   `json:"price" db:"price" validate:"required"`
	TaxItem   services.TaxItemInterface `json:"tax_item" db:"-"`
}

// String is not required by pop and may be deleted
func (t Tax) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Taxes is not required by pop and may be deleted
type Taxes []Tax

// String is not required by pop and may be deleted
func (t Taxes) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Tax) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Name, Name: "Name"},
		&validators.StringIsPresent{Field: t.TaxCode, Name: "TaxCode"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Tax) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Tax) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// AfterFind bind current Tax object with TaxItem
func (t *Tax) AfterFind(tx *pop.Connection) error {
	ti := services.NewTaxItem(t.TaxCode, t.Price)
	ti.Calculate()

	t.TaxItem = ti
	return nil
}
