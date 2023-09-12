package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:] // Exclude the program name from the arguments

	// Check if the filename is missing
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	// Check if too many arguments are provided
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	// Read the file
	fileName := args[0]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Display the file content
	fmt.Print(string(content))
}
