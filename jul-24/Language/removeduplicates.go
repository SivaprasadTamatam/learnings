package main

import "fmt"

func removeDuplicates(s string) string {
	stack := []rune{}

	for _, ch := range s {
		if len(stack) > 0 && stack[len(stack)-1] == ch {
			fmt.Println(string(stack))
			stack = stack[:len(stack)-1]
			fmt.Println(string(stack))
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}

func main() {
	example1 := "abbaca"
	//example2 := "azxxzy"

	fmt.Println("Output for example 1:", removeDuplicates(example1)) // Output: "ca"
	//fmt.Println("Output for example 2:", removeDuplicates(example2)) // Output: "ay"
}
