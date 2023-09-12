package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	for i, ch := range s {
		if i == 0 && (ch == '+' || ch == '-') {
			if ch == '-' {
				sign = -1
			}
			continue
		}
		if ch < '0' || ch > '9' {
			return 0
		}
		digit := int(ch - '0')
		result = result*10 + digit
	}
	return sign * result
}
