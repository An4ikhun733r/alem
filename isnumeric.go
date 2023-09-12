package piscine

func IsNumeric(s string) bool {
	for index, letter := range s {
		index++
		if letter < 48 || letter > 57 {
			return false
		}
	}
	return true
}
