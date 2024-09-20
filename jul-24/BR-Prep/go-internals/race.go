package main

import (
	"fmt"
	"sync"
)

// Mutex

var (
	counter int
	mu      sync.Mutex
)

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}
