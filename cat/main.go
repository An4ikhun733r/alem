package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// Read from standard input (stdin)
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			os.Stderr.WriteString("Error reading from stdin: " + err.Error() + "\n")
			os.Exit(1) // Exit with status 1
		}
	} else {
		// Read and print each file
		for _, arg := range args {
			file, err := os.Open(arg)
			if err != nil {
				os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			defer file.Close()

			_, err = io.Copy(os.Stdout, file)
			if err != nil {
				os.Stderr.WriteString("Error reading file " + ": " + err.Error() + "\n")
				os.Exit(1)
			}
		}
	}

	os.Exit(0) // Exit with status 0
}
