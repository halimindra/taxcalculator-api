package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTobaccoTax(t *testing.T) {
	tt := NewTobaccoTax(10000)
	tt.Calculate()

	assert.Equal(t, tt.Type, TypeTobacco)
	assert.NotNil(t, tt.Refundable)
	assert.NotZero(t, tt.Price)
	assert.NotZero(t, tt.TaxAmount)
	assert.NotZero(t, tt.Subtotal)
	t.Logf("%+v", tt)
}
