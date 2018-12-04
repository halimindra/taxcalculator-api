package services

// constants
const (
	TypeTobacco       = "Tobacco"
	TypeFood          = "Food & Beverage"
	TypeEntertainment = "Entertainment"
)

// TaxItemInterface is interface for TaxItem
type TaxItemInterface interface {
	Calculate()
}

// TaxItem base struct for Tax object
type TaxItem struct {
	Code       string  `json:"-"`
	Type       string  `json:"type"`
	Refundable bool    `json:"refundable"`
	Price      float64 `json:"-"`
	TaxAmount  float64 `json:"tax_amount"`
	Subtotal   float64 `json:"subtotal"`
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
