package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChangeBFSPractice(coins []int, amout int) int {
	if amout == 0 {
		return 0
	}

	queue := []int{0}
	minCois := 0
	v := make(map[int]bool)
	v[0] = true

	for len(queue) > 0 {
		minCois++
		nextQueue := []int{}
		for _, current := range queue {
			for _, coin := range coins {
				next := current + coin
				if next == amout {
					return minCois
				}
				if next < amout && v[next] == false {
					nextQueue = append(nextQueue, next)
					v[next] = true
				}

			}
		}
		queue = nextQueue

	}
	return -1
}

func coinChangeBFS(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	queue := []int{0}
	minCoins := 0

	v := make(map[int]bool)

	v[0] = true
	for len(queue) > 0 {
		minCoins++
		nextQueue := []int{}

		for _, current := range queue {
			for _, coin := range coins {
				next := current + coin
				if next == amount {
					fmt.Println(next, current, coin)
					return minCoins
				}

				if next < amount && !v[next] {
					nextQueue = append(nextQueue, next)
					v[next] = true
				}
			}
		}
		fmt.Println(nextQueue)
		queue = nextQueue
	}
	return -1
}

func main() {
	coins1 := []int{1, 2, 3}
	amount1 := 11

	fmt.Println("Example 1:", coinChangeBFS(coins1, amount1)) // Output: 3 (11 = 5 + 5 + 1)

}
