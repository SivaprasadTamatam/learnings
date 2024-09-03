package main

import "fmt"

func longPressChar(name, typed string) bool {
	i, j := 0, 0

	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return i == len(name)
}

func main() {
	fmt.Println(longPressChar("leelee", "lleeelee"))
}
