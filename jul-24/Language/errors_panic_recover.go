package main

import "fmt"

func causePanic() {
	fmt.Println("About to panic")
	panic("something went wrong")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoved from panic : ", r)
		}
	}()
	causePanic()

	fmt.Println("this will be not printed ")
}

func main() {
	fmt.Println("starting main")
	safeFunction()
	fmt.Println("Completed")
}
