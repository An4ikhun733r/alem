package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		os.Exit(0) // Exit with status 0 if no arguments are provided
	}

	// Read the value of -c option
	var numBytes int
	if args[0] == "-c" && len(args) >= 2 {
		numBytes = parseInt(args[1])
		if numBytes < 0 {
			fmt.Printf("Invalid value for -c option: %s\n", args[1])
			os.Exit(1) // Exit with status 1
		}
		args = args[2:]
	}

	// Process each file
	exitCode := 0 // Initialize exit code as 0
	for i, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			exitCode = 1 // Update exit code to 1
			continue
		}
		defer file.Close()

		if i > 0 {
			fmt.Println()
		}

		if len(args) > 1 {
			fmt.Printf("==> %s <==\n", filename)
		}

		_, err = tail(file, numBytes)
		if err != nil {
			fmt.Printf("Error reading file %s\n", err.Error())
			exitCode = 1 // Update exit code to 1
		}
	}

	os.Exit(exitCode) // Exit with the determined exit code
}

func parseInt(s string) int {
	n := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		} else {
			return -1
		}
	}
	return n
}

func tail(file *os.File, numBytes int) ([]byte, error) {
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := stat.Size()
	if int64(numBytes) > fileSize {
		numBytes = int(fileSize)
	}

	offset := fileSize - int64(numBytes)
	_, err = file.Seek(offset, 0) // Seek to the beginning of the file relative to the current position
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, numBytes)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	_, err = os.Stdout.Write(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
