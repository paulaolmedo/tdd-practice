package rest_api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestEmptyCart(t *testing.T) {
	cart := Cart{}
	assert.Empty(t, cart.Books)
}

func TestNotEmptyCart(t *testing.T){
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
		Editorial: Editorial{Library: []string{book1.Isbn, book2.Isbn} },
	}

	cart.AddBook(book1)
	cart.AddBook(book2)
	assert.NotEmpty(t, cart.Books)
	assert.Equal(t, 2, len(cart.Books))
}

func TestShouldNotAddExternalBooks(t *testing.T){
	editorial := Editorial{Library: []string{"isbn1", "isbn2"}}
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
	assert.NotEmpty(t, editorial.Library)

}