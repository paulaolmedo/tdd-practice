package rest_api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyCart(t *testing.T) {
	cart := NewCart()
	assert.Empty(t, cart.Books)
}

func TestNotEmptyCart(t *testing.T){
	cart := NewCart()
	cart.AddBook("isbn1")
	cart.AddBook("isbn2")
	assert.NotEmpty(t, cart.Books)
	assert.Equal(t, 2, len(cart.Books))
}

func TestShouldNotAddExternalBooks(t *testing.T){
	cart := NewCart()
	result := cart.AddBook("isbn0")
	assert.False(t, result)
	assert.Empty(t, cart.Books)
}

func TestAddSameBookShouldIncrementQuantity(t *testing.T)  {
	cart := NewCart()
	cart.AddBook("isbn1")
	cart.AddBook("isbn1")
	assert.NotEmpty(t, cart.Books)
	assert.Equal(t, 2, cart.Books["isbn1"], "Books quantity must be 2 ")

}
