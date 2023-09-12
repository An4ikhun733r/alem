package piscine

func Join(strs []string, sep string) string {
	var result string

	for i, str := range strs {
		if i > 0 {
			result += sep
		}
		result += str
	}

	return result
}
