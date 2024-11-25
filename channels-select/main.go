package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	evenNumbers := make(chan int)
	oddNumbers := make(chan int)
	errChan := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 30; i++ {
			if i > 20 {
				errChan <- errors.New(fmt.Sprintf("error: number %d is greater than 20", i))
			} else if i%2 == 0 {
				evenNumbers <- i
			} else {
				oddNumbers <- i
			}
		}
		close(evenNumbers)
		close(oddNumbers)
	}()

	for {
		select {
		case num, ok := <-evenNumbers:
			if ok {
				fmt.Printf("Received an even number: %d\n", num)
			}
		case num, ok := <-oddNumbers:
			if ok {
				fmt.Printf("Received an odd number: %d\n", num)
			}
		case err, ok := <-errChan:
			if ok {
				fmt.Println(err)
			}
		case <-time.After(1 * time.Second):
			wg.Wait()
			return
		}
	}
}
