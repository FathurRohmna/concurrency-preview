package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	evenNumbers := make(chan int)
	oddNumbers := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 20; i++ {
			if i%2 == 0 {
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
		case <-time.After(1 * time.Second):
			wg.Wait()
			return
		}
	}

}
