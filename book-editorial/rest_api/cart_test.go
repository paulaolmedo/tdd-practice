package rest_api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestValidClientID(t *testing.T) {
	var bookService BookStoreService
	ok, err := bookService.ValidateUser("client1", "password1")

	require.NoError(t, err)
	assert.True(t, ok)
}

func TestInvalidClientID(t *testing.T) {
	var bookService BookStoreService
	notOk, err := bookService.ValidateUser("clientID", "password")

	assert.Error(t, err, "el cliente no existe")
	assert.False(t, notOk)
}

func TestInvalidPassword(t *testing.T) {
	var bookService BookStoreService
	notOk, err := bookService.ValidateUser("client1", "someotherpassword")

	assert.Error(t, err, "la contraseña es inválida")
	assert.False(t, notOk)
}

func TestCreateCartWithValidUser(t *testing.T) {
	var service BookStoreService

	id, err := service.CreateCart("client1", "password1")
	require.NoError(t, err)

	cart, err := service.FindCart(id)
	require.NoError(t, err)

	assert.NotEmpty(t, cart)
}

func TestFindInvalidCart(t *testing.T) {
	var service BookStoreService

	cart, err := service.FindCart("someotherid")

	assert.Error(t, err, "no such cart")
	assert.Empty(t, cart)
}

func TestAddToCart(t *testing.T) {
	var service BookStoreService

	id, err := service.CreateCart("client1", "password1")
	require.NoError(t, err)

	cart, err := service.FindCart(id)
	require.NoError(t, err)

	ok := cart.AddBook("isbn1")
	expectedLength := 1
	actualLength := len(cart.Books)

	assert.True(t, ok)
	assert.NotEmpty(t, cart.Books)
	assert.Equal(t, expectedLength, actualLength)
}
