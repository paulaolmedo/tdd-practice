package roman

import "errors"

func ConvertToRomanString(romanNumber int) (string, error) {
	if romanNumber <= 0 {
		return "", errors.New("cannot be negative")
	}

	return "expectedArabic", nil
}
