package piscine

func StrRev(s string) string {
	a := []rune(s)
	for i := 0; i < len(s); i++ {
		a[i] = rune(s[len(s)-i-1])
	}
	return string(a)
}
