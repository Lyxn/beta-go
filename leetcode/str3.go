package leetcode

import (
	"fmt"
	"math"
	"strings"
)

func expandCenter(s string, lo, hi int) (lo1, hi1 int) {
	for lo >= 0 && hi < len(s) && s[lo] == s[hi] {
		lo--
		hi++
	}
	return lo + 1, hi - 1
}

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	maxLen := -1
	maxLo := 0
	maxHi := 0
	for i := 0; i < n-1; i++ {
		lo1, hi1 := expandCenter(s, i, i)
		if hi1-lo1 > maxLen {
			maxLen = hi1 - lo1
			maxHi = hi1
			maxLo = lo1
		}
		if s[i] != s[i+1] {
			continue
		}
		lo2, hi2 := expandCenter(s, i, i+1)
		if hi2-lo2 > maxLen {
			maxLen = hi2 - lo2
			maxHi = hi2
			maxLo = lo2
		}
	}
	return s[maxLo : maxHi+1]
}

func minCut(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i < n; i++ {
		l0, r0 := expandCenter(s, i, i)
		for l0 <= r0 {
			dp[r0+1] = min(dp[r0+1], dp[l0]+1)
			l0++
			r0--
		}
		if s[i-1] == s[i] {
			l1, r1 := expandCenter(s, i-1, i)
			for l1 < r1 {
				dp[r1+1] = min(dp[r1+1], dp[l1]+1)
				l1++
				r1--
			}
		}
	}
	return dp[n] - 1
}

func countSubstrings(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	expand := func(l, r int) int {
		res := 0
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			res++
		}
		return res
	}
	cnt := 1
	for i := 1; i < n; i++ {
		cnt += expand(i, i)
		if s[i] == s[i-1] {
			cnt += expand(i-1, i)
		}
	}
	return cnt
}

func longestDecomposition(text string) int {
	n := len(text)
	if n <= 1 {
		return n
	}
	for i := 1; i < n/2+1; i++ {
		if text[0:i] == text[n-i:n] {
			return 2 + longestDecomposition(text[i:n-i])
		}
	}
	return 1
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				if i == j-1 {
					dp[i][j] = 2
				} else {
					dp[i][j] = 2 + dp[i+1][j-1]
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func strangePrinter(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	ss := make([]byte, 0)
	ss = append(ss, s[0])
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ss = append(ss, s[i])
		}
	}

	ns := len(ss)
	dp := make([][]int, ns)
	for i := 0; i < ns; i++ {
		dp[i] = make([]int, ns)
		dp[i][i] = 1
	}
	for i := ns - 2; i >= 0; i-- {
		for j := i + 1; j < ns; j++ {
			dp[i][j] = 1 + dp[i+1][j]
			for k := i + 1; k <= j; k++ {
				if ss[i] != ss[k] {
					continue
				}
				if k+1 < ns {
					dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k+1][j])
				} else {
					dp[i][j] = min(dp[i][j], dp[i][k-1])
				}
			}
		}
	}
	return dp[0][ns-1]
}

func repeatedSubString(s string) int {
	ss := s + s
	res := strings.Index(ss[1:len(ss)-1], s)
	return res
}

func encode(s string) string {
	n := len(s)
	dp := make([][]string, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]string, n)
		dp[i][i] = s[i : i+1]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if i+4 > j {
				dp[i][j] = s[i : j+1]
				continue
			}
			p := repeatedSubString(s[i : j+1])
			tmp := s[i : j+1]
			if p != -1 {
				cnt := (j + 1 - i) / (p + 1)
				tmp = fmt.Sprintf("%d[%s]", cnt, dp[i][i+p])
			}
			for k := i; k < j; k++ {
				tc := dp[i][k] + dp[k+1][j]
				if len(tc) < len(tmp) {
					tmp = tc
				}
			}
			dp[i][j] = tmp
		}
	}
	return dp[0][n-1]
}
