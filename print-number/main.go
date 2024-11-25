package main

import (
	"fmt"
)

// function to print number
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// function to print letter
func printLetters() {
	for letter := 'a'; letter <= 'j'; letter++ {
		fmt.Printf("%c\n", letter)
	}
}

func main() {
	// Run printNumbers and printLetters
	printNumbers()
	printLetters()
}
