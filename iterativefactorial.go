package piscine

func IterativeFactorial(nb int) int {
	if nb >= 0 && nb < 21 {
		result := 1
		for i := 1; i <= nb; i++ {
			result = result * i
		}
		return result
	}
	return 0
}
