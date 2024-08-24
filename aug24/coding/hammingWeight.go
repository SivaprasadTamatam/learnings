package main

import "fmt"

func HammingWeight(num uint32) int {
	res := 0

	for num != 0 {
		res += int(num & 1)
		num >>= 1
	}
	return res
}

func main() {
	fmt.Println(HammingWeight(uint32(3)))
}
