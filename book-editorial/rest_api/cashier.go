package rest_api

import (
	"errors"
	"strconv"
	"time"
)

type Cashier struct {
}

type CreditCard struct {
	Number          int64
	ExpirationMonth string
	ExpirationYear  string
	CardOwner string
}

func NewCashier() Cashier {
	return Cashier{}
}

func (c Cashier) ValidateCreditCard(card *CreditCard) (interface{}, error) {
	if card == nil {
		return "", errors.New("credit card can not be empty")
	}

	if recursionCountDigits(card.Number) != 16 {
		return "", errors.New("invalid credit card number")
	}

	month, err := strconv.Atoi(card.ExpirationMonth)
	if err != nil {
		return "", errors.New("invalid expiration month")

	}

	if len(card.ExpirationMonth) != 2 || !(month > 0 && month < 13) {
		return "", errors.New("invalid expiration month")
	}

	year, err := strconv.Atoi(card.ExpirationYear)
	if err != nil {
		return "", errors.New("invalid expiration year")

	}

	if len(card.ExpirationYear) != 4 {
		return "", errors.New("invalid expiration year")
	}

	now := time.Now()
	cardExpirationDate := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.Local)

	if cardExpirationDate.Before(now) {
		return "", errors.New("expired credit card")
	}

	if card.CardOwner == "" || len(card.CardOwner) > 30{
		return "", errors.New("invalid card owner")
	}

	return "", nil
}

func (c Cashier) ProcessCart(cart Cart, card CreditCard) (interface{}, error) {
	if len(cart.Books) == 0 {
		return nil, errors.New("cart is empty")
	}
	return nil, nil
}

func recursionCountDigits(number int64) int64 {
	if number < 10 {
		return 1
	} else {
		return 1 + recursionCountDigits(number/10)
	}
}
