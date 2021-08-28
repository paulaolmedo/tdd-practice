package roman

import "errors"

func ConvertToRomanString(arabic int) (string, error) {
	var romanNumber string

	if arabic <= 0 {
		return "", errors.New("cannot be negative")
	} else if arabic >= 10 && arabic < 20 {
		romanNumber += "X"
		arabic = arabic - 10
	} else if arabic >= 20 {
		romanNumber += "XX"
		arabic = arabic - 20
	}

	if arabic <= 3 {
		for i := 0; i < arabic; i++ {
			romanNumber += "I"
		}
	} else if arabic == 4 {
		romanNumber += "IV"
	} else if arabic <= 8 {
		romanNumber += "V"
		for i := 5; i < arabic; i++ {
			romanNumber += "I"
		}
	} else {
		romanNumber += "IX"
	}

	return romanNumber, nil
}
