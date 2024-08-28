package main

func CoinChange(coins []int, amount int) int {

	if amount == 0 || len(coins) == 0 {
		return 0
	}

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
