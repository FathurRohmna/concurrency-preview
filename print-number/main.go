package main

import (
	"fmt"
	"sync"
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
	var wg sync.WaitGroup

	// Run printNumbers and printLetters concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		printNumbers()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		printLetters()
	}()

	// // Set time.Sleep to prevent the program exit before the goroutines complete their task
	// time.Sleep(10 * time.Second)

	wg.Wait()
}
