package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/halimindra/taxcalculator-api/models"
	"github.com/halimindra/taxcalculator-api/utils"
)

// TaxesResource is the resource for the Tax model
type TaxesResource struct {
	buffalo.Resource
}

// List gets all Taxes. This function is mapped to the path
// GET /taxes
func (v TaxesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(
			http.StatusInternalServerError,
			r.JSON(utils.NewErrorResponse("No transaction found")),
		)
	}

	taxes := &models.Taxes{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Taxes from the DB
	if err := q.All(taxes); err != nil {
		return c.Render(
			http.StatusNoContent,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, taxes))
}

// Show gets the data for one Tax. This function is mapped to
// the path GET /taxes/{tax_id}
func (v TaxesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(
			http.StatusInternalServerError,
			r.JSON(utils.NewErrorResponse("No transaction found")),
		)
	}

	// Allocate an empty Tax
	tax := &models.Tax{}

	// To find the Tax the parameter tax_id is used.
	if err := tx.Find(tax, c.Param("tax_id")); err != nil {
		return c.Render(
			http.StatusNotFound,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	return c.Render(200, r.Auto(c, tax))
}

// Create adds a Tax to the DB. This function is mapped to the
// path POST /taxes
func (v TaxesResource) Create(c buffalo.Context) error {
	// Allocate an empty Tax
	tax := &models.Tax{}

	// Bind request to model
	if err := c.Bind(tax); err != nil {
		return c.Render(
			http.StatusBadRequest,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(
			http.StatusInternalServerError,
			r.JSON(utils.NewErrorResponse("No transaction found")),
		)
	}

	// Validate the data
	_, err := tx.ValidateAndCreate(tax)
	if err != nil {
		return c.Render(
			http.StatusUnprocessableEntity,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	// and redirect to the taxes index page
	return c.Render(201, r.Auto(c, tax))
}

// Update changes a Tax in the DB. This function is mapped to
// the path PUT /taxes/{tax_id}
func (v TaxesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(
			http.StatusInternalServerError,
			r.JSON(utils.NewErrorResponse("No transaction found")),
		)
	}

	// Allocate an empty Tax
	tax := &models.Tax{}

	if err := tx.Find(tax, c.Param("tax_id")); err != nil {
		return c.Render(
			http.StatusNotFound,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	// Bind request to model
	if err := c.Bind(tax); err != nil {
		return c.Render(
			http.StatusBadRequest,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	_, err := tx.ValidateAndUpdate(tax)
	if err != nil {
		return c.Render(
			http.StatusUnprocessableEntity,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	// and redirect to the taxes index page
	return c.Render(200, r.Auto(c, tax))
}

// Destroy deletes a Tax from the DB. This function is mapped
// to the path DELETE /taxes/{tax_id}
func (v TaxesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(
			http.StatusInternalServerError,
			r.JSON(utils.NewErrorResponse("No transaction found")),
		)
	}

	// Allocate an empty Tax
	tax := &models.Tax{}

	// To find the Tax the parameter tax_id is used.
	if err := tx.Find(tax, c.Param("tax_id")); err != nil {
		return c.Render(
			http.StatusNotFound,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	if err := tx.Destroy(tax); err != nil {
		return c.Render(
			http.StatusUnprocessableEntity,
			r.JSON(utils.NewErrorResponse(err.Error())),
		)
	}

	// Redirect to the taxes index page
	return c.Render(200, r.Auto(c, tax))
}
