package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	counts map[string]int
	mux    *sync.RWMutex
}

func (s *safeCounter) Inc(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.counts[key]++
}

func (s *safeCounter) Get(key string) int {

	return s.counts[key]
}

func main() {
	s := &safeCounter{counts: make(map[string]int), mux: &sync.RWMutex{}}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Inc("Key")
		}()
	}

	for i := 0; i < 10; i++ {

		fmt.Println(s.Get("Key"))

	}

	wg.Wait()

}
