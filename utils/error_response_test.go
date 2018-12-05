package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewErrorResponse(t *testing.T) {
	er := NewErrorResponse("Unknown error")
	assert.NotNil(t, er)
}
