package leetcode

import "math"

func findTheLongestSubstring(s string) int {
	n := len(s)
	numState := 32
	dp := make([]int, numState)
	dp[0] = -1
	for i := 1; i < numState; i++ {
		dp[i] = math.MaxInt32
	}
	st := byte(0)
	maxLen := 0
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			st = st ^ 1
		} else if s[i] == 'e' {
			st = st ^ (1 << 1)
		} else if s[i] == 'i' {
			st = st ^ (1 << 2)
		} else if s[i] == 'o' {
			st = st ^ (1 << 3)
		} else if s[i] == 'u' {
			st = st ^ (1 << 4)
		}
		if dp[st] == math.MaxInt32 {
			dp[st] = i
		} else {
			l := i - dp[st]
			if l > maxLen {
				maxLen = l
			}
		}
	}
	return maxLen
}
