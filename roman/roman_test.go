package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNumber(t *testing.T) {
	arabic1 := 1
	arabic8 := 8

	roman1, err1 := ConvertToRomanString(arabic1)
	roman8, err8 := ConvertToRomanString(arabic8)

	assert.Empty(t, err1)
	assert.Equal(t, "I", roman1)

	assert.Empty(t, err8)
	assert.Equal(t, "VIII", roman8)
}

func TestInvalidNumber(t *testing.T) {
	arabic := 0

	roman, err := ConvertToRomanString(arabic)

	expectedRoman := ""
	expectedError := "cannot be negative"

	assert.Equal(t, expectedRoman, roman)
	assert.EqualError(t, err, expectedError)
}
