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

	tokenTypes := map[byte]string{
		'(': "LEFT_PAREN",
		')': "RIGHT_PAREN",
		'{': "LEFT_BRACE",
		'}': "RIGHT_BRACE",
		',': "COMMA",
		'.': "DOT",
		'-': "MINUS",
		'+': "PLUS",
		';': "SEMICOLON",
		'*': "STAR",
		'=': "EQUAL",
		'!': "BANG",
		'<': "LESS",
		'>': "GREATER",
		'/': "SLASH",
	}
	containsLexicalError := false
	lineNumber := 1
	// Scanner implementation
	for i := 0; i < len(fileContents); i++ {
		c := fileContents[i]

		// Ignore whitespace characters
		if c == ' ' || c == '\t' || c == '\n' {
			if c == '\n' {
				lineNumber++
			}
			continue
		}
		// Go implicitly converts char literal to corresponding ascii code so it compares numbers

		// Handle two-character lexemes
		if i+1 < len(fileContents) {
			nextChar := fileContents[i+1]
			switch c {
			case '=':
				if nextChar == '=' {
					fmt.Println("EQUAL_EQUAL == null")
					i++ // Consume next character
					continue
				}
			case '!':
				if nextChar == '=' {
					fmt.Println("BANG_EQUAL != null")
					i++ // Consume next character
					continue
				}
			case '<':
				if nextChar == '=' {
					fmt.Println("LESS_EQUAL <= null")
					i++ // Consume next character
					continue
				}
			case '>':
				if nextChar == '=' {
					fmt.Println("GREATER_EQUAL >= null")
					i++ // Consume next character
					continue
				}
			case '/':
				if nextChar == '/' {
					// Skip to the next newline character
					for i < len(fileContents) && fileContents[i] != '\n' {
						i++
					}
					lineNumber++
					continue
				} else {
					fmt.Println("SLASH / null")
					continue // Consume the '/' character
				}
			}
		}
		// Handle single-character lexemes
		if tokenType, ok := tokenTypes[c]; ok {
			fmt.Printf("%s %c null\n", tokenType, c)
		} else {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", lineNumber, c)
			containsLexicalError = true
		}
	}
	fmt.Println("EOF  null")
	if containsLexicalError {
		os.Exit(65)
	} else {
		os.Exit(0)
	}
}
