package models

import (
	"testing"
)

func Test_Tax(t *testing.T) {
	tax := Tax{
		Name:    "Item 1",
		TaxCode: "1",
		Price:   999999,
	}

	if err := DB.Create(&tax); err != nil {
		t.Error(err)
	}
}
