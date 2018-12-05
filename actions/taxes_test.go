package actions

import (
	"strconv"

	"github.com/halimindra/taxcalculator-api/models"
)

func (as *ActionSuite) Test_TaxesResource_List() {
	res := as.JSON("/taxes").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_TaxesResource_Show() {
	tax := models.Tax{
		Name:    "Test Tax Show",
		TaxCode: "1",
		Price:   10000,
	}
	res := as.JSON("/taxes").Post(tax)
	as.Equal(201, res.Code)

	// retreive the first tax from the database
	err := as.DB.First(&tax)
	as.NoError(err)
	as.NotZero(tax.ID)

	res = as.JSON("/taxes/" + strconv.Itoa(tax.ID)).Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_TaxesResource_Create() {
	tax := models.Tax{
		Name:    "Test Tax Create",
		TaxCode: "2",
		Price:   10000,
	}
	res := as.JSON("/taxes").Post(tax)
	as.Equal(201, res.Code)

	// retreive the first tax from the database
	err := as.DB.First(&tax)
	as.NoError(err)
	as.NotZero(tax.ID)

	// assert the tax name was saved correctly
	as.Equal("Test Tax Create", tax.Name)
}

func (as *ActionSuite) Test_TaxesResource_Update() {
	tax := models.Tax{
		Name:    "Test Tax Update",
		TaxCode: "3",
		Price:   10000,
	}
	updatedTax := models.Tax{}
	res := as.JSON("/taxes").Post(tax)
	as.Equal(201, res.Code)

	// retreive the first tax from the database
	err := as.DB.First(&tax)
	as.NoError(err)
	as.NotZero(tax.ID)

	// do update request
	tax.Name = "Test Tax Updated"
	res = as.JSON("/taxes/" + strconv.Itoa(tax.ID)).Put(&tax)
	as.Equal(200, res.Code)

	// retrieve the updated tax from the database
	err = as.DB.Find(&updatedTax, tax.ID)
	as.Equal(updatedTax.ID, tax.ID)
	as.Equal(updatedTax.Name, "Test Tax Updated")
}

func (as *ActionSuite) Test_TaxesResource_Destroy() {
	tax := models.Tax{
		Name:    "Test Tax Destroy",
		TaxCode: "3",
		Price:   10000,
	}
	res := as.JSON("/taxes").Post(tax)
	as.Equal(201, res.Code)

	// retreive the first tax from the database
	err := as.DB.First(&tax)
	as.NoError(err)
	as.NotZero(tax.ID)

	// do delete request
	res = as.JSON("/taxes/" + strconv.Itoa(tax.ID)).Delete()
	as.Equal(200, res.Code)

	// check if it's deleted from database
	err = as.DB.Find(&tax, tax.ID)
	as.NotNil(err)
}
