package main

import (
	"fmt"
	"sync"
	"time"
)

// function to print number
func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() 
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

// function to print letter
func printLetters(wg *sync.WaitGroup) {
	defer wg.Done() 
	for letter := 'a'; letter <= 'j'; letter++ {
		fmt.Printf("%c\n", letter)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	// Run printNumbers and printLetters concurrently
	go printNumbers(&wg)
	go printLetters(&wg)

	// // Set time.Sleep to prevent the program exit before the goroutines complete their task
	// time.Sleep(10 * time.Second)

	wg.Wait()
}
