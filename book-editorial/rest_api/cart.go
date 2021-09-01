package rest_api

import (
	"errors"
	"github.com/google/uuid"
)

type Book struct {
	Isbn string
	Price float32
}

type Cart struct {
	Id string
	Books []Book
	TotalAmount float32
	Editorial Editorial
}

func NewCart(client, password string) (Cart, error) {
	if client == "" || password == ""{
		return Cart{}, errors.New("1 | Client or Password can not be empty")
	}

	id, err := uuid.NewUUID()
	if err != nil{
		return Cart{}, errors.New("1 | Can not generate id")
	}

	return Cart{
		Id:          id.String(),
		Books:       nil,
		TotalAmount: 0,
		Editorial:   Editorial{},
	}, nil
}

func (c *Cart) AddBook(book Book) bool{
	for _, isbn := range c.Editorial.Library {
		if isbn == book.Isbn {
			c.Books = append(c.Books, book)
			return true
		}
	}
	return false
}



