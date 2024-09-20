package main

import "fmt"

func main() {
	a := 10
	b := 20

	c := new(int)
	*c = 30

	fmt.Println(a, b, *c)
}
