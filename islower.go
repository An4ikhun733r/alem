package piscine

func IsLower(s string) bool {
	for index, letter := range s {
		index++
		if letter < 97 || letter > 122 {
			return false
		}
	}
	return true
}
