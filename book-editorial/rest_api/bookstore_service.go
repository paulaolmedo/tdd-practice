package rest_api

import (
	"errors"
	"time"
)

// registeredClients representa los clientes registrados en el sistema
var registeredClients = map[string]string{
	"client1": "password1",
	"client2": "password2",
}

// createdCarts representa los carritos creados por el cliente
var createdCarts = map[string]Cart{
	"cartId": Cart{},
}

type BookStoreService struct {
	update time.Time
}

func (b *BookStoreService) ValidateUser(inputClientID string, inputPassword string) (bool, error) {
	if password, ok := registeredClients[inputClientID]; !ok {
		return false, errors.New("el cliente no existe")
	} else if inputPassword != password {
		return false, errors.New("la contraseña es inválida")
	}

	return true, nil
}

func (b *BookStoreService) CreateCart(clientID string, password string) (string, error) {
	b.update = time.Now()

	_, err := b.ValidateUser(clientID, password)
	if err != nil {
		return "", err
	}

	cart := NewCart(clientID)
	createdCarts[cart.Id] = cart

	return cart.Id, nil
}

func (b *BookStoreService) FindCart(id string) (Cart, error) {
	thirtyMin := time.Now().Add(time.Minute * -30)
	if b.update.Before(thirtyMin) {
		return Cart{}, errors.New("no such cart")
	}

	if cart, ok := createdCarts[id]; ok {
		b.update = time.Now()
		return cart, nil
	} else {
		return Cart{}, errors.New("no such cart")
	}
}

func (b *BookStoreService) AddBook(id string, book string) error {
	cart, _ := b.FindCart(id)

	thirtyMin := time.Now().Add(time.Minute * -30) // 11:30
	if b.update.Before(thirtyMin) {
		return errors.New("invalid cart date") // ver
	}

	ok := cart.AddBook(book)
	if !ok {
		return errors.New("error")
	}
	b.update = time.Now()
	return nil
}
