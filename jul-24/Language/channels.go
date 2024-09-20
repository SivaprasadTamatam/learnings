package main

import "fmt"

func main() {
	ch := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//close(ch)
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

}
