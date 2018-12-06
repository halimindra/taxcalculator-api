package services

// FoodTax inherit TaxItem
type FoodTax struct {
	TaxItem
}

// NewFoodTax create new FoodTax instance
func NewFoodTax(price float64) *FoodTax {
	return &FoodTax{
		TaxItem{
			Code:         "1",
			Type:         TypeFood,
			Refundable:   true,
			Price:        price,
			IsCalculated: false,
		},
	}
}

// Calculate implementation of custom logic for tax calculation
func (ft *FoodTax) Calculate() {
	ft.TaxAmount = 0.1 * ft.Price
	ft.Subtotal = ft.Price + ft.TaxAmount
	ft.IsCalculated = true
}
