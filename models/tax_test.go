package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tax(t *testing.T) {
	tax := Tax{
		Name:    "Item 1",
		TaxCode: "1",
		Price:   999999,
	}

	err := DB.Create(&tax)
	assert.Nil(t, err)
}
