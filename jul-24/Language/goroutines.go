package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	cn := 1
	wg.Add(100)

	for i := 1; i <= 100; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			fmt.Println(cn)
			cn++
		}()
	}

	wg.Wait()
}
