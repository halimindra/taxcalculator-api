package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewEntertainmentTax(t *testing.T) {
	et := NewEntertainmentTax(10000)
	et.Calculate()

	assert.Equal(t, et.Type, TypeEntertainment)
	assert.NotNil(t, et.Refundable)
	assert.NotZero(t, et.Price)
	assert.NotZero(t, et.TaxAmount)
	assert.NotZero(t, et.Subtotal)
	t.Logf("%+v", et)
}
