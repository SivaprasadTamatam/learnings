package main

import "fmt"

func Heapify(nums []int, n, i int) {
	largetst := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && nums[left] > nums[largetst] {
		largetst = left
	}

	if right < n && nums[right] > nums[largetst] {
		largetst = right
	}

	if largetst != i {
		nums[i], nums[largetst] = nums[largetst], nums[i]
		Heapify(nums, n, largetst)
	}
}

func BuildMaxHeap(nums []int) {
	n := len(nums)
	st := (n / 2) - 1
	for i := st; i >= 0; i-- {
		Heapify(nums, n, i)
	}
}
func ExtractMax(nums []int, n int) int {

	if n == 0 {
		return -1
	}
	res := nums[0]
	nums[0], nums[n-1] = nums[n-1], nums[0]
	nums = nums[:n-1]
	Heapify(nums, n-1, 0)
	return res
}
func main() {
	nums := []int{3, 19, 1, 14, 8, 7}
	fmt.Println(nums)
	BuildMaxHeap(nums)
	fmt.Println(nums)

	for i := 0; i <= len(nums); i++ {
		fmt.Println(ExtractMax(nums, len(nums)-i))
		//fmt.Println(nums)
	}
	fmt.Println(nums)
}
