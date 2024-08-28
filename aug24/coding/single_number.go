package main

import "fmt"

func SingleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

func main() {
	fmt.Println(SingleNumber([]int{1, 2, 2, 3, 4, 1, 4}))
}
