package piscine

func Compact(ptr *[]string) int {
	cnt := 0
	var result []string
	for _, ch := range *ptr {
		if ch != "" {
			result = append(result, ch)
			cnt++
		}
	}
	*ptr = result
	return cnt
}
