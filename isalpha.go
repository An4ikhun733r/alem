package piscine

func IsAlpha(s string) bool {
	for index, letter := range s {
		index++
		if (letter < 97 || letter > 122) && (letter < 65 || letter > 90) && (letter < 48 || letter > 57) {
			return false
		}
	}
	return true
}
