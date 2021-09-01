package rest_api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEditorialNotEmptyLibrary(t *testing.T){
	editorial := Editorial{
		Library: []string{"1234"},
	}
	assert.NotEmpty(t, editorial.Library)
}