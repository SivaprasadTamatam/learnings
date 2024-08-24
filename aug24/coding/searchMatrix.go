package main

import "fmt"

func SearchMatrix(m [][]int, target int) bool {
	if len(m) == 0 || len(m[0]) == 0 {
		return false
	}

	row, col := 0, len(m[0])-1

	for row < len(m) && col >= 0 {
		if m[row][col] == target {
			return true
		} else if m[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 61
	output := SearchMatrix(matrix, target)
	fmt.Println("Is the target in the matrix?:", output)
}
