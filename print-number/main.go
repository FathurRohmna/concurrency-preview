package main

import (
	"fmt"
	"time"
)

// function to print number
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

// function to print letter
func printLetters() {
	for letter := 'a'; letter <= 'j'; letter++ {
		fmt.Printf("%c\n", letter)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Run printNumbers and printLetters concurrently
	go printNumbers()   
	go printLetters()

	// Set time.Sleep to prevent the program exit before the goroutines complete their task
	time.Sleep(10 * time.Second)
}
