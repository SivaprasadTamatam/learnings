package main

import (
	"fmt"
	"sync"
)

func Even(max int, ec, oc chan int, wg *sync.WaitGroup) {
	for i := 0; i < max; i += 2 {
		ec <- i
		<-oc
	}

	wg.Done()
}

func Odd(max int, ec, oc chan int, wg *sync.WaitGroup) {
	for i := 1; i <= max; i += 2 {
		fmt.Println(<-ec)
		fmt.Println(i)
		oc <- i
	}
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	ec, oc := make(chan int), make(chan int)
	wg.Add(2)
	go Even(10, ec, oc, &wg)
	go Odd(10, ec, oc, &wg)

	wg.Wait()
}
