/*
Problem: Given a string, find the length of the longest substring without repeating characters.
Approach: Use a sliding window approach with a map to track characters and their indices.
*/
package main

import "fmt"

func LongestSubstring(s string) int {
	mp := make(map[byte]int)
	st, ml := 0, 0

	for i := 0; i < len(s); i++ {
		if idx, ok := mp[s[i]]; ok && idx >= st {
			st = idx + 1
			fmt.Println(st)
		}

		mp[s[i]] = i

		if (i - st + 1) > ml {
			ml = i - st + 1
		}
	}
	r := []rune(s)
	fmt.Println(string(r[st : st+ml]))
	return ml
}

func main() {
	fmt.Println(LongestSubstring("acbdabcde"))
}
