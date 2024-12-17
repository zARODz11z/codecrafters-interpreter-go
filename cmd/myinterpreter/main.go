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

	reservedWords := map[string]string{
		"and":    "AND",
		"class":  "CLASS",
		"else":   "ELSE",
		"false":  "FALSE",
		"for":    "FOR",
		"fun":    "FUN",
		"if":     "IF",
		"nil":    "NIL",
		"or":     "OR",
		"print":  "PRINT",
		"return": "RETURN",
		"super":  "SUPER",
		"this":   "THIS",
		"true":   "TRUE",
		"var":    "VAR",
		"while":  "WHILE",
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

		if c == '"' {
			// Handle string literals
			start := i
			i++ // Move past the opening quote
			for i < len(fileContents) && fileContents[i] != '"' {
				i++
			}
			if i >= len(fileContents) {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", lineNumber)
				containsLexicalError = true
				break
			}
			// i is now at the closing quote
			fmt.Printf("STRING \"%s\" %s\n", fileContents[start+1:i], fileContents[start+1:i])
			continue
		}

		if isAlpha(c) {
			start := i
			for i < len(fileContents) && (isAlpha(fileContents[i]) || isDigit(fileContents[i])) {
				i++
			}
			lexeme := string(fileContents[start:i])
			if tokenType, ok := reservedWords[lexeme]; ok {
				fmt.Printf("%s %s null\n", tokenType, lexeme)
			} else {
				fmt.Printf("IDENTIFIER %s null\n", lexeme)
			}
			i-- // Decrement i to reprocess the non-alphanumeric character
			continue
		}

		if c >= '0' && c <= '9' {
			start := i
			containsDecimal := false
			i++ // Move past the first digit
			for i < len(fileContents) && (fileContents[i] >= '0' && fileContents[i] <= '9') {
				i++
			}
			if i < len(fileContents) && fileContents[i] == '.' {
				containsDecimal = true
				i++ // Move past the decimal point

				for i < len(fileContents) && (fileContents[i] >= '0' && fileContents[i] <= '9') {
					i++
				}
			}
			numberLexeme := string(fileContents[start:i])
			numberLiteral := numberLexeme
			// If the number has a decimal point, remove unecessary trailing zeros
			if containsDecimal {
				numberLiteral = trimTrailingZeros(numberLiteral)
			}
			// For Whole numbers, add ".0" to match the books format
			if !containsDecimal {
				numberLiteral += ".0"
			}
			fmt.Printf("NUMBER %s %s\n", numberLexeme, numberLiteral)
			i-- // Decrement i to reprocess the non-digit character
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

func trimTrailingZeros(number string) string {
	// Remove trailing zeros after the decimal point.
	// Ensure that at least one digit remains after the decimal.
	for len(number) > 1 && number[len(number)-1] == '0' && number[len(number)-2] != '.' {
		number = number[:len(number)-1] // Trim the last character (zero).
	}
	// If the last character is a '.', remove it as well.
	// This handles cases like "34." -> "34".
	if len(number) > 1 && number[len(number)-1] == '.' {
		number = number[:len(number)-1] // Trim the decimal point.
	}
	return number // Return the properly formatted number string.
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
