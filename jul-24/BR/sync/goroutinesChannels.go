package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	ch := make(chan bool, 1)
	increment := func() {
		defer wg.Done()

		<-ch
		counter++
		ch <- true
	}

	ch <- true
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()

	fmt.Println(counter)
}
