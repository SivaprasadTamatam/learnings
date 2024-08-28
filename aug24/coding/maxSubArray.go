package main

import (
	"fmt"
	"math"
)

func MaxSubArray(nums []int) (int, int, int) {
	ms := math.MinInt32
	cs, start, s, e := 0, 0, 0, 0

	for i, num := range nums {
		if num > cs+num {
			cs = num
			s = i
		} else {
			cs = cs + num
		}

		if cs > ms {
			ms = cs
			start = s
			e = i
		}
	}
	return ms, start, e
}

func main() {
	fmt.Println(MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
