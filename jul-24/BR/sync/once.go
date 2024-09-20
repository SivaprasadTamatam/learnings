package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	var wg sync.WaitGroup

	initiliaze := func() {
		fmt.Println("Initializing")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initiliaze)
		}()

	}
	wg.Wait()
}
