package services

// TobaccoTax inherit TaxItem
type TobaccoTax struct {
	TaxItem
}

// NewTobaccoTax create new TobaccoTax instance
func NewTobaccoTax(price float64) *TobaccoTax {
	return &TobaccoTax{
		TaxItem{
			Code:       "2",
			Type:       TypeTobacco,
			Refundable: false,
			Price:      price,
		},
	}
}

// Calculate get Calculated tax amount
func (tt *TobaccoTax) Calculate() {
	tt.TaxAmount = 10 + (0.02 * tt.Price)
	tt.Subtotal = tt.Price + tt.TaxAmount
}
