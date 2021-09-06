package rest_api

import "errors"

var registeredClients = map[string]string{
	"client1": "password1",
	"client2": "password2",
}

var createdCarts = map[string]Cart{
	"cartId": Cart{},
}

type BookStoreService struct {
	
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
	_, err := b.ValidateUser(clientID, password)
	if err != nil {
		return "", err
	}

	cart := NewCart(clientID)
	createdCarts[cart.Id] = cart

	return cart.Id, nil
}

func (b *BookStoreService) FindCart(id string) (Cart, error) {
	if cart, ok := createdCarts[id]; ok {
		return cart, nil
	} else {
		return Cart{}, errors.New("no such cart")
	}
}
