package services

// constants
const (
	TypeTobacco       = "Tobacco"
	TypeFood          = "Food & Beverage"
	TypeEntertainment = "Entertainment"
)

// TaxItemInterface is interface for TaxItem
type TaxItemInterface interface {
	GetTaxAmount() float64
	Calculate()
}

// TaxItem base struct for Tax object
type TaxItem struct {
	Code         string  `json:"-"`
	Type         string  `json:"type"`
	Refundable   bool    `json:"refundable"`
	Price        float64 `json:"-"`
	TaxAmount    float64 `json:"tax_amount"`
	Subtotal     float64 `json:"subtotal"`
	IsCalculated bool    `json:"-"`
}

// NewTaxItem create new TaxItem instance
func NewTaxItem(code string, price float64) TaxItemInterface {
	switch code {
	case "1":
		return NewFoodTax(price)
	case "2":
		return NewTobaccoTax(price)
	default:
		return NewEntertainmentTax(price)
	}
}

// Calculate get Calculated tax amount
func (ti *TaxItem) Calculate() {}

// GetTaxAmount retrieve TaxAmount
func (ti *TaxItem) GetTaxAmount() float64 {
	if !ti.IsCalculated {
		ti.Calculate()
	}
	return ti.TaxAmount
}
