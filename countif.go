package piscine

func CountIf(f func(string) bool, tab []string) int {
	cnt := 0
	for _, ch := range tab {
		if f(ch) == true {
			cnt++
		}
	}
	return cnt
}
