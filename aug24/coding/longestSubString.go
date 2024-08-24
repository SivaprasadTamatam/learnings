package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	n, ml, left := len(s), 0, 0
	for right := 0; right < n; right++ {
		if v, ok := mp[s[right]]; ok && v >= left {
			left = v + 1
		}

		mp[s[right]] = right
		if ml < (right - left + 1) {
			ml = right - left + 1
		}
	}
	return ml
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
