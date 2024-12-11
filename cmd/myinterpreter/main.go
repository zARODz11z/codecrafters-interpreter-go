package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	//Uncomment this block to pass the first stage

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Scanner implementation
	for i := 0; i < len(fileContents); i++ {
		c := fileContents[i]
		// Go implicitly converts char literal to corresponding ascii code so it compares numbers
		switch c {
		case '(':
			fmt.Println("LEFT_PAREN ( null")
		case ')':
			fmt.Println("RIGHT_PAREN ) null")
		}
	}

	fmt.Println("EOF  null")

}
