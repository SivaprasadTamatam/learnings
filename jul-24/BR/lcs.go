package main

import "fmt"

func longestCommonSubsequence(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println("Example 1:", longestCommonSubsequence(text1, text2)) // Output: 3

	text1 = "abc"
	text2 = "abc"
	fmt.Println("Example 2:", longestCommonSubsequence(text1, text2)) // Output: 3

	text1 = "abc"
	text2 = "def"
	fmt.Println("Example 3:", longestCommonSubsequence(text1, text2)) // Output: 0
}
