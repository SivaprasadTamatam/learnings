package main

import (
	"fmt"
	"sync"
)

func main() {
	var rwm sync.RWMutex
	var counter int

	readCounter := func(wg *sync.WaitGroup) {
		defer wg.Done()
		rwm.RLock()
		fmt.Println("Read - ", counter)
		rwm.RUnlock()
	}

	incrementCOunter := func(wg *sync.WaitGroup) {
		defer wg.Done()
		rwm.Lock()
		counter++
		rwm.Unlock()
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readCounter(&wg)

		wg.Add(1)
		go incrementCOunter(&wg)
	}

	wg.Wait()

	fmt.Println(counter)

}
