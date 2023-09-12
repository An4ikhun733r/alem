package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	path := os.Args[0]
	programName := 0
	for i := 0; i < len(os.Args[0]); i++ {
		if path[i] == '/' {
			programName = i + 1
		}
	}
	name := os.Args[0][programName:]
	for _, ch := range name {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
