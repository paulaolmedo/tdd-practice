package rest_api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyCart(t *testing.T) {
	cart := NewCart()
	assert.Empty(t, cart.Books)
}

func TestNotEmptyCart(t *testing.T) {
	cart := NewCart()
	cart.AddBook("isbn1")
	cart.AddBook("isbn2")

	actualLength := len(cart.Books)
	expectedLength := 2

	assert.NotEmpty(t, cart.Books)
	assert.Equal(t, expectedLength, actualLength)
}

func TestShouldNotAddExternalBooks(t *testing.T) {
	cart := NewCart()
	result := cart.AddBook("isbn0")

	assert.False(t, result)
	assert.Empty(t, cart.Books)
}

func TestAddSameBookShouldIncrementQuantity(t *testing.T) {
	cart := NewCart()
	cart.AddBook("isbn1")
	cart.AddBook("isbn1")

	actualLength := cart.Books["isbn1"]
	expectedLength := 2

	assert.NotEmpty(t, cart.Books)
	assert.Equal(t, expectedLength, actualLength, "Books quantity must be 2 ")
}



