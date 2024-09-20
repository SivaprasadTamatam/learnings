package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan struct{})

	go func() {
		mu.Lock()
		defer mu.Unlock()
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("Oeration Completed")
	case <-ctx.Done():
		fmt.Println("Timedout")
	}
}
