package services

// EntertainmentTax inherit TaxItem
type EntertainmentTax struct {
	TaxItem
}

// NewEntertainmentTax create new EntertainmentTax instance
func NewEntertainmentTax(price float64) *EntertainmentTax {
	return &EntertainmentTax{
		TaxItem{
			Code:         "3",
			Type:         TypeEntertainment,
			Refundable:   false,
			Price:        price,
			IsCalculated: false,
		},
	}
}

// Calculate get Calculated tax amount
func (et *EntertainmentTax) Calculate() {
	var taxAmount float64
	if et.Price >= 100 {
		taxAmount = 0.01 * (et.Price - 100)
	}
	et.TaxAmount = taxAmount
	et.Subtotal = et.Price + et.TaxAmount
	et.IsCalculated = true
}
