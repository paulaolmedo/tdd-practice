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
	cart := Cart{
		Id:          "1",
		Books:       []Book{},
		TotalAmount: 0,
	}
	book1 := Book{
		Isbn:  "1234567890876",
		Price: 120,
	}
	book2 := Book{
		Isbn:  "1234567890876",
		Price: 120,
	}
	cart.AddBook(book1)
	cart.AddBook(book2)
	assert.NotEmpty(t, cart.Books)
	assert.Equal(t, 2, len(cart.Books))
}

func TestShouldNotAddExternalBooks(t *testing.T){

}