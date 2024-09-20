/*
Question 3

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.



Example 1:

Input: coins = [1,2,5], amount = 11

Output: 3

Explanation: 11 = 5 + 5 + 1

Example 2:

Input: coins = [2], amount = 3

Output: -1

Example 3:

Input: coins = [1], amount = 0

Output: 0
*/

package main

import "fmt"

func coinChangeDP(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if (i - coin) >= 0 {
				dp[i] = min(dp[i], 1+dp[i-coin])
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

func changeCoins(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	queue := []int{0}
	mp := make(map[int]bool)
	mp[0] = true
	minCoins := 0

	for len(queue) > 0 {
		minCoins++
		newQueue := []int{}

		for _, current := range queue {
			for _, coin := range coins {
				next := current + coin

				if next == amount {
					return minCoins
				}

				if next < amount && !mp[next] {
					newQueue = append(newQueue, next)
					mp[next] = true
				}
			}

		}
		queue = newQueue
	}
	return -1
}

func main() {
	fmt.Println(changeCoins([]int{1, 2, 5}, 11))
	fmt.Println(changeCoins([]int{2}, 3))
	fmt.Println(changeCoins([]int{1}, 0))
	fmt.Println(coinChangeDP([]int{1, 2, 5}, 11))
	fmt.Println(coinChangeDP([]int{2}, 3))
	fmt.Println(coinChangeDP([]int{1}, 0))
}
