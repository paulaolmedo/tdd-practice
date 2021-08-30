package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNumber(t *testing.T) {
	arabic1 := 1
	roman1, err1 := ConvertToRomanString(arabic1)

	arabic9 := 9
	roman9, err9 := ConvertToRomanString(arabic9)

	assert.Empty(t, err1)
	assert.Equal(t, "I", roman1)

	assert.Empty(t, err9)
	assert.Equal(t, "IX", roman9)
}

func TestInvalidNumber(t *testing.T) {
	arabic := 0

	roman, err := ConvertToRomanString(arabic)

	expectedRoman := ""
	expectedError := "cannot be negative"

	assert.Equal(t, expectedRoman, roman)
	assert.EqualError(t, err, expectedError)
}

func TestMulti2(t *testing.T) {
	arabic10 := 10
	roman10, err10 := ConvertToRomanString(arabic10)

	arabic19 := 19
	roman19, err19 := ConvertToRomanString(arabic19)

	assert.Empty(t, err10)
	assert.Equal(t, "X", roman10)

	assert.Empty(t, err19)
	assert.Equal(t, "XIX", roman19)
}

func TestMulti3(t *testing.T) {
	arabic20 := 20
	roman20, err20 := ConvertToRomanString(arabic20)

	arabic29 := 29
	roman29, err29 := ConvertToRomanString(arabic29)

	assert.Empty(t, err20)
	assert.Equal(t, "XX", roman20)

	assert.Empty(t, err29)
	assert.Equal(t, "XXIX", roman29)
}

func TestMulti4(t *testing.T) {
	arabic20 := 30
	roman20, err20 := ConvertToRomanString(arabic20)

	arabic29 := 39
	roman29, err29 := ConvertToRomanString(arabic29)

	assert.Empty(t, err20)
	assert.Equal(t, "XXX", roman20)

	assert.Empty(t, err29)
	assert.Equal(t, "XXXIX", roman29)
}