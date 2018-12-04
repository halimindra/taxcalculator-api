package grifts

import (
	"github.com/halimindra/taxcalculator-api/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
