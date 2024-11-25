package main

import (
	"fmt"
	"sync"
)

func produce(result chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println("Producing:", i)
		result <- fmt.Sprintf("%d", i)
	}
	close(result)
}

func consume(result chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range result {
		fmt.Println(num)
		fmt.Println("Consuming:", num)
	}
}

func main() {
	var wg sync.WaitGroup

	// Buffered: Transaction processes can be performed simultaneously according to the set capacity
	// Unbuffered: Transaction processes will be done one by one, waiting for each process to finish before proceeding to the next
	result := make(chan string)

	wg.Add(2)
	go produce(result, &wg)
	go consume(result, &wg)
	

	wg.Wait()
}
