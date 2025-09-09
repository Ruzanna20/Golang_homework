package main

import (
	"fmt"
	"sync"
)

// count of goroutine
const workerCount = 3

func checkNumbers(idworker int, numbers <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numbers {
		if num%2 == 0 {
			fmt.Printf("Worker %v: %v zuyg e\n", idworker, num)
		} else {
			fmt.Printf("Worker %v: %v kent e\n", idworker, num)
		}

	}
}

func main() {
	numbers := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go checkNumbers(i, numbers, &wg)
	}

	nums := []int{2, 3, 6, 7, 10}
	go func() {
		for _, num := range nums {
			numbers <- num
		}
		close(numbers)
	}()

	wg.Wait()
}

