package main

import (
	"fmt"
	"sort"
)

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	results := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			results = append(results, left[i])
			i++
		} else {
			results = append(results, right[j])
			j++
		}
	}

	if i < len(left) {
		results = append(results, left[i:]...)

	}

	if j < len(right) {
		results = append(results, right[j:]...)
	}
	return results
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}

	fmt.Println("Original array:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)

	sort.Ints(arr)
	fmt.Println("Sorted array:", sortedArr)
}
