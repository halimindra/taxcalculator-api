package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/halimindra/taxcalculator-api/models"
	"github.com/halimindra/taxcalculator-api/utils"
)

// BillShow default implementation.
func BillShow(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(
			http.StatusInternalServerError,
			r.JSON(utils.NewErrorResponse("No transaction found")),
		)
	}

	// Allocate an empty Taxes
	taxes := models.Taxes{}

	// Retrieve all Taxes from the DB
	if err := tx.All(&taxes); err != nil {
		return c.Render(
			http.StatusNoContent,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	// Create Bill from Taxes
	bill, err := models.NewBill(taxes)
	if err != nil {
		return c.Render(
			http.StatusNoContent,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	return c.Render(200, r.Auto(c, bill))
}
