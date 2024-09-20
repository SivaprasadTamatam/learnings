package main

import "fmt"

func Sum[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1.3, 4.6))
}
