package piscine

func LastRune(s string) rune {
	for index, letter := range s {
		if index == len(s)-1 {
			return letter
		}
	}
	return 0
}
