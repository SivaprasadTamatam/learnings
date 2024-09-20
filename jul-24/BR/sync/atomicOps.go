package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	increment := func(wg *sync.WaitGroup) {
		defer wg.Done()
		atomic.AddInt64(&counter, 1)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	fmt.Println(counter)
}
