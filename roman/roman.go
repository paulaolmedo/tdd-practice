package roman

import "errors"

func ConvertToRomanString(arabic int) (string, error) {
	var romanNumber string

	if arabic <= 0 {
		return "", errors.New("cannot be negative")
	}

	tens := arabic / 10
	units := arabic % 10

	if tens <= 3 {
		for i := 0; i < tens; i++ {
			romanNumber += "X"
		}
	} else if tens == 4 {
		romanNumber += "XL"
	} else if tens <= 8 {
		romanNumber += "L"
		for i := 5; i < tens; i++ {
			romanNumber += "X"
		}
	} else {
		romanNumber += "XC"
	}

	if units <= 3 {
		for i := 0; i < units; i++ {
			romanNumber += "I"
		}
	} else if units == 4 {
		romanNumber += "IV"
	} else if units <= 8 {
		romanNumber += "V"
		for i := 5; i < units; i++ {
			romanNumber += "I"
		}
	} else {
		romanNumber += "IX"
	}

	return romanNumber, nil
}
