package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := len(arg) - 1; i > 0; i-- {
		for _, ch := range arg[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
