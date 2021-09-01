package rest_api

import (
	"github.com/google/uuid"
)

var catalog = []string{"isbn1", "isbn2", "isbn3", "isbn4"}

type Cart struct {
	Id          string
	Catalog     []string
	Books       map[string]int
	TotalAmount float32
}

func NewCart() Cart {
	id, _ := uuid.NewUUID()
	cart := Cart{
		Id:      id.String(),
		Catalog: catalog,
		Books:   map[string]int{},
	}
	return cart

}

func (c *Cart) AddBook(book string) bool {
	for _, isbn := range c.Catalog {
		if isbn == book {
			if quantity, ok := c.Books[isbn]; ok {
				c.Books[isbn] = quantity + 1
			}
			c.Books[isbn] = 1
			return true
		}
	}
	return false
}
