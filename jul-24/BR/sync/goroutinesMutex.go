package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var counter int

	increment := func(wg *sync.WaitGroup) {
		defer wg.Done()
		mu.Lock()
		fmt.Println(counter)
		counter++
		mu.Unlock()
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	fmt.Println(counter)

}
