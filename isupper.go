package piscine

func IsUpper(s string) bool {
	for index, letter := range s {
		index++
		if letter < 65 || letter > 90 {
			return false
		}
	}
	return true
}
