package piscine

func IsPrintable(s string) bool {
	for index, letter := range s {
		index++
		if letter == '\n' || letter == '\r' || letter == '\t' || letter == '\b' {
			return false
		}
	}
	return true
}
