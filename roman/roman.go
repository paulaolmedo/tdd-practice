package roman

import "errors"

func ConvertToRomanString(arabic int) (string, error) {
	if arabic <= 0 {
		return "", errors.New("cannot be negative")
	}
	var romanNumber string
	if arabic <= 3 {
		for i := 0; i < arabic; i++ {
			romanNumber += "I"
		}
	} else if arabic == 4 {
		romanNumber = "IV"
	} else {
		romanNumber += "V"
		for i := 5; i < arabic; i++ {
			romanNumber += "I"
		}
	}

	return romanNumber, nil
}
