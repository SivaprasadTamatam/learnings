/*
Question 2

Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.



Example 1:

Input: text1 = "abcde", text2 = "ace"

Output: 3

Explanation: The longest common subsequence is "ace" and its length is 3.

Example 2:

Input: text1 = "abc", text2 = "abc"

Output: 3

Explanation: The longest common subsequence is "abc" and its length is 3.

Example 3:

Input: text1 = "abc", text2 = "def"

Output: 0

Explanation: There is no such common subsequence, so the result is 0.
*/

package main

import "fmt"

func LCS(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
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
	fmt.Println(LCS("abcde", "ace"))
	fmt.Println(LCS("abc", "abc"))
	fmt.Println(LCS("abc", "def"))
}
