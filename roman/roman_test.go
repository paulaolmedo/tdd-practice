package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNumber(t *testing.T) {
	roman := 1

	arabic, _ := ConvertToRomanString(roman)

	assert.Equal(t, "expectedArabic", arabic)
}

func TestInvalidNumber(t *testing.T) {
	roman := 0

	arabic, err := ConvertToRomanString(roman)

	expectedArabic := ""
	expectedError := "cannot be negative"

	assert.Equal(t, expectedArabic, arabic)
	assert.EqualError(t, err, expectedError)
}
