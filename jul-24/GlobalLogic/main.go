package main

import (
	"fmt"
)

func EvenOrOdd(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	var n int
	for {
		fmt.Println("Please Enter the Number: (Type 100 if you want Quit)")
		fmt.Scanf("%d", &n)
		fmt.Println(n)
		if n == 100 {
			break
		}

		fmt.Println(EvenOrOdd(n))
	}

}
