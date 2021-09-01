package rest_api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyCart(t *testing.T) {
	cart := Cart{}
	assert.Empty(t, cart.Books)
}

func TestNotEmptyCart(t *testing.T){
	editorial := NewEditorial()

	book1 := Book{
		Isbn:  "isbn1",
		Price: 120,
	}
	book2 := Book{
		Isbn:  "isbn2",
		Price: 120,
	}

	cart := Cart{
		Id:          "1",
		Books:       []Book{},
		TotalAmount: 0,
		Editorial: editorial,
	}

	cart.AddBook(book1)
	cart.AddBook(book2)
	assert.NotEmpty(t, cart.Books)
	assert.Equal(t, 2, len(cart.Books))
}

func TestShouldNotAddExternalBooks(t *testing.T){
	cart := Cart{
		Id:          "1",
		Books:       []Book{},
		TotalAmount: 0,
	}
	book := Book{
		Isbn:  "isbn3",
		Price: 120,
	}
	result := cart.AddBook(book)
	assert.False(t, result)
	assert.Empty(t, cart.Books)

}
//para la interfaz
func TestCreateInvalidCart(t *testing.T) {
	cart, err := NewCart("", "pass")
	assert.Empty(t, cart)
	assert.Error(t, err)
	assert.EqualError(t, err, "1 | Client or Password can not be empty")
}

func TestCreateCart(t *testing.T) {
	cart, err := NewCart("mario", "pass")
	assert.NoError(t, err)
	assert.NotEmpty(t, cart.Id)
}