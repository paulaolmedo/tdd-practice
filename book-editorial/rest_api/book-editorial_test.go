package rest_api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEditorialNotEmptyLibrary(t *testing.T){
	editorial := NewEditorial()
	assert.NotEmpty(t, editorial.Library)
}