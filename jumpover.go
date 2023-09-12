package piscine

func JumpOver(str string) string {
	newline := "\n"
	result := ""
	if len(str) < 3 {
		return newline
	}
	for i, ch := range str {
		if (i+1)%3 == 0 {
			result += string(ch)
		}
	}
	result += string("\n")
	return result
}
