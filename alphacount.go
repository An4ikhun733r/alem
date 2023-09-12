package piscine

func AlphaCount(s string) int {
	count := 0
	for index, letter := range s {
		index++
		if (letter >= 97 && letter <= 122) || (letter >= 65 && letter <= 90) {
			count = count + 1
		}
	}
	return count
}
