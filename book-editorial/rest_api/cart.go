package rest_api

type Book struct {
	Isbn string
	Price float32
}

type Cart struct {
	Id string
	Books []Book
	TotalAmount float32
}

func (c Cart) AddBook(book Book) {
	c.Books = append(c.Books, book)
}



