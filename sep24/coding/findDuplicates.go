package main

import (
	"fmt"
	"math"
)

func findDuplicates(nums []int) []int {
	res := []int{}

	for _, num := range nums {
		index := int(math.Abs(float64(num)) - 1)
		fmt.Println(index)
		if nums[index] < 0 {
			res = append(res, int(math.Abs(float64(num))))
			fmt.Println(res)
		} else {
			nums[index] = -nums[index]
		}
		fmt.Println(nums)
	}
	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}
