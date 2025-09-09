package main

import (
	"fmt"
	"sync"
)

func loopingOnNames(m map[string]int, mutex *sync.Mutex) {
	names := []string{"Anna", "Ani", "Aram", "Aren"}

	var wg sync.WaitGroup

	for i, name := range names {
		wg.Add(1)
		go func(key string, val int) {
			defer wg.Done()
			updateMap(key, val, m, mutex)
		}(name, i)
	}
	wg.Wait()
}

func updateMap(key string, val int, m map[string]int, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	m[key] = val
}

func main() {
	m := make(map[string]int)
	var mutex sync.Mutex
	//Synchronizing goroutines with WaitGroup
	var wg sync.WaitGroup

	// Wait for 1 goroutine to finisհ
	wg.Add(1)
	go func() {
		defer wg.Done()
		loopingOnNames(m, &mutex)
	}()

	// Wait for gorountines to finish
	wg.Wait()

	mutex.Lock()
	fmt.Println(m)
	mutex.Unlock()
}

