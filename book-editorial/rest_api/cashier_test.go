package rest_api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateCashier(t *testing.T) {
	cashier := NewCashier()
	assert.NotNil(t, cashier)
}

func TestShouldValidateEmptyCreditCard(t *testing.T) {
	cashier := NewCashier()
	_, err := cashier.ValidateCreditCard(nil)
	assert.Error(t, err)
	assert.EqualError(t, err, "credit card can not be empty")
}

func TestInvalidCreditCardNumber(t *testing.T) {
	cashier := NewCashier()
	creditCard := CreditCard{
		Number: 231,
	}
	_, err := cashier.ValidateCreditCard(&creditCard)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid credit card number")
}

func TestValidCreditCardNumber(t *testing.T) {
	cashier := NewCashier()
	creditCard := CreditCard{
		Number:          4349008516656431,
		ExpirationMonth: "01",
		ExpirationYear:  "9999",
		CardOwner: "PAULA",
	}
	_, err := cashier.ValidateCreditCard(&creditCard)
	assert.NoError(t, err)
}

func TestInvalidExpirationMonthFormat(t *testing.T) {
	cashier := NewCashier()
	creditCard := CreditCard{
		Number:          4349008516656431,
		ExpirationMonth: "1",
	}
	_, err := cashier.ValidateCreditCard(&creditCard)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid expiration month")
}

func TestInvalidExpirationYearFormat(t *testing.T) {
	cashier := NewCashier()
	creditCard := CreditCard{
		Number:          4349008516656431,
		ExpirationMonth: "10",
		ExpirationYear:  "11",
	}
	_, err := cashier.ValidateCreditCard(&creditCard)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid expiration year")
}

func TestExpiredCreditCard(t *testing.T) {
	cashier := NewCashier()
	creditCard := CreditCard{
		Number:          4349008516656431,
		ExpirationMonth: "01",
		ExpirationYear:  "1980",
	}
	_, err := cashier.ValidateCreditCard(&creditCard)
	assert.Error(t, err)
	assert.EqualError(t, err, "expired credit card")
}

func TestValidCard(t *testing.T) {
	cashier := NewCashier()
	creditCard := CreditCard{
		Number:          4349008516656431,
		ExpirationMonth: "01",
		ExpirationYear:  "9999",
		CardOwner: "MARIO THOMAS",
	}
	_, err := cashier.ValidateCreditCard(&creditCard)
	assert.NoError(t, err)
}

func TestInvalidCardOwner(t *testing.T) {
	cashier := NewCashier()
	creditCard := CreditCard{
		Number:          4349008516656431,
		ExpirationMonth: "01",
		ExpirationYear:  "9999",
		CardOwner: "VILLA CARLOS PAZ CORDOBA ARGENTINA",
	}
	_, err := cashier.ValidateCreditCard(&creditCard)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid card owner")
}

func TestEmptyCardOwner(t *testing.T) {
	cashier := NewCashier()
	creditCard := CreditCard{
		Number:          4349008516656431,
		ExpirationMonth: "01",
		ExpirationYear:  "9999",
		CardOwner: "",
	}
	_, err := cashier.ValidateCreditCard(&creditCard)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid card owner")
}

func TestProcessCart(t *testing.T) {
	cashier := NewCashier()
	cart := NewCart()
	cart.AddBook("isbn1")
	creditCard := CreditCard{
		Number:          4349008516656431,
		ExpirationMonth: "01",
		ExpirationYear:  "9999",
		CardOwner: "PECA",
	}
	_, err := cashier.ProcessCart(cart, creditCard)
	assert.NoError(t, err)
}

func TestProcessEmptyCart(t *testing.T) {
	cashier := NewCashier()
	cart := NewCart()
	creditCard := CreditCard{
		Number:          4349008516656431,
		ExpirationMonth: "01",
		ExpirationYear:  "2024",
	}
	_, err := cashier.ProcessCart(cart, creditCard)
	assert.Error(t, err)
	assert.EqualError(t, err, "cart is empty")
}