package piscine

func digitToValue(digit byte) int {
	switch {
	case digit >= '0' && digit <= '9':
		return int(digit - '0')
	case digit >= 'A' && digit <= 'Z':
		return int(digit - 'A' + 10)
	case digit >= 'a' && digit <= 'z':
		return int(digit - 'a' + 10)
	default:
		return -1
	}
}

func valueToDigit(value int) byte {
	if value >= 0 && value <= 9 {
		return byte(value + '0')
	}
	return byte(value - 10 + 'A')
}

func convertToDecimal(nbr, baseFrom string) int {
	base := len(baseFrom)
	decimal := 0

	for i := 0; i < len(nbr); i++ {
		digit := nbr[i]
		value := digitToValue(digit)
		decimal = decimal*base + value
	}

	return decimal
}

func convertFromDecimal(decimal int, baseTo string) string {
	base := len(baseTo)
	result := ""

	for decimal > 0 {
		remainder := decimal % base
		digit := valueToDigit(remainder)
		result = string(digit) + result
		decimal = decimal / base
	}

	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := convertToDecimal(nbr, baseFrom)
	result := convertFromDecimal(decimal, baseTo)
	return result
}
