package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for l := '9'; l >= '0'; l-- {
				for m := '9'; m >= '0'; m-- {
					if (i > l) || (i == l && j > m) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(l)
						z01.PrintRune(m)
						if i == '0' && j == '1' && l == '0' && m == '0' {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
