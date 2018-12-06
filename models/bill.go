package models

type Bill struct {
	PriceSubtotal float64 `json:"price_subtotal" db:"-"`
	TaxSubtotal   float64 `json:"tax_subtotal" db:"-"`
	GrandTotal    float64 `json:"grand_subtotal" db:"-"`
	Taxes         []Tax   `json:"taxes" db:"-"`
}

// NewBill create new Bill from Taxes
func NewBill(taxes Taxes) (Bill, error) {
	var priceSubtotal float64
	var taxSubtotal float64
	var grandTotal float64

	for _, tax := range taxes {
		priceSubtotal += tax.Price
		taxSubtotal += tax.TaxItem.GetTaxAmount()
	}

	grandTotal = priceSubtotal + taxSubtotal
	b := Bill{
		PriceSubtotal: priceSubtotal,
		TaxSubtotal:   taxSubtotal,
		GrandTotal:    grandTotal,
		Taxes:         taxes,
	}

	return b, nil
}
