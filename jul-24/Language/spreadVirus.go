package main

import (
	"fmt"
)

func spreadVirus(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	directions := [][]int{
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
		{-1, 0}, // up
	}

	queue := make([][2]int, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	days := 0
	for len(queue) > 0 {
		newQueue := make([][2]int, 0)
		for _, pos := range queue {
			for _, dir := range directions {
				newRow, newCol := pos[0]+dir[0], pos[1]+dir[1]
				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					newQueue = append(newQueue, [2]int{newRow, newCol})
				}
			}
		}
		if len(newQueue) == 0 {
			break
		}
		queue = newQueue
		days++
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				return -1 // If there is any house that is not affected by the virus
			}
		}
	}

	return days
}

func main() {
	grid := [][]int{
		{0, 1, 2},
		{1, 1, 0},
		{0, 1, 1},
	}

	days := spreadVirus(grid)
	if days == -1 {
		fmt.Println("Not all houses can be affected by the virus.")
	} else {
		fmt.Printf("It will take %d days for the virus to spread across all houses.\n", days)
	}
}
