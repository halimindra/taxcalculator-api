package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewFoodTax(t *testing.T) {
	ft := NewFoodTax(10000)
	ft.Calculate()

	assert.Equal(t, ft.Type, TypeFood)
	assert.NotNil(t, ft.Refundable)
	assert.NotZero(t, ft.Price)
	assert.NotZero(t, ft.TaxAmount)
	assert.NotZero(t, ft.Subtotal)
	t.Logf("%+v", ft)
}
