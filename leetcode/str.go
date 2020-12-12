package leetcode

import "math"

func isValidParen(sr string) bool {
	var stk []rune
	for i := 0; i < len(sr); i++ {
		if sr[i] == '(' {
			stk = append(stk, ')')
		} else if sr[i] == '{' {
			stk = append(stk, '}')
		} else if sr[i] == '[' {
			stk = append(stk, ']')
		} else if sr[i] == ')' || sr[i] == '}' || sr[i] == ']' {
			num := len(stk)
			if num == 0 {
				return false
			} else if stk[num-1] != rune(sr[i]) {
				return false
			}
			stk = stk[:num-1]
		}
	}
	return len(stk) == 0
}

func EditDist(word1 string, word2 string) (ret int) {
	n := len(word1)
	m := len(word2)
	mat := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		mat[i] = make([]int, m+1)
		mat[i][0] = i
	}
	for j := 0; j <= m; j++ {
		mat[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				mat[i][j] = mat[i-1][j-1]
			} else {
				mat[i][j] = MinInt(mat[i-1][j], mat[i][j-1], mat[i-1][j-1]) + 1
			}
		}
	}
	return mat[n][m]
}

func LengthOfLongestSubstring(s string) (ret int) {
	n := len(s)
	hm := make(map[byte]int, n)
	if len(s) == 0 {
		return 0
	}
	start := 0
	win := 0
	maxLen := 0
	for i := 0; i < n; i++ {
		c := s[i]
		t, ok := hm[c]
		if !ok {
			win += 1
			hm[c] = i
		} else if t < start {
			win += 1
			hm[c] = i
		} else {
			if win > maxLen {
				maxLen = win
			}
			start = t + 1
			win = i - start + 1
			hm[c] = i
		}
	}
	if win > maxLen {
		maxLen = win
	}
	return maxLen
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

func expandCenter(s string, lo, hi int) (lo1, hi1 int) {
	for lo >= 0 && hi < len(s) && s[lo] == s[hi] {
		lo--
		hi++
	}
	return lo + 1, hi - 1
}

func longestDupSubstring(S string) string {
	n := len(S)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = int(S[i] - 'a')
	}
	l := 1
	r := n
	for l < r {
		m := (l + r + 1) / 2
		x := searchDup(nums, m)
		if x >= 0 {
			l = m
		} else {
			r = m - 1
		}
	}
	x := searchDup(nums, l)
	if x != -1 {
		return S[x : x+l]
	}
	return ""
}

func searchDup(nums []int, l int) int {
	n := len(nums)
	h := 0
	rdx := 26
	mod := math.MaxInt32
	hl := powMod(rdx, l-1, mod)
	for i := 0; i < l; i++ {
		h = (h*rdx + nums[i]) % mod
	}
	dct := make(map[int][]int, n-l)
	dct[h] = []int{0}
	for i := 1; i < n-l+1; i++ {
		h = ((h-nums[i-1]*hl)*rdx + nums[i+l-1]) % mod
		if h < 0 {
			h += mod
		}
		ss, ok := dct[h]
		if ok {
			for _, st := range ss {
				if checkSame(nums, st, i, l) {
					return i
				}
			}
			dct[h] = append(ss, i)
		} else {
			dct[h] = []int{i}
		}
	}
	return -1
}

func checkSame(nums []int, i, j, l int) bool {
	for k := 0; k < l; k++ {
		if nums[i+k] != nums[j+k] {
			return false
		}
	}
	return true
}

func powMod(b, p, mod int) int {
	if p == 0 {
		return 1
	} else if p == 1 {
		return b % mod
	}
	if p%2 == 0 {
		b1 := powMod(b, p/2, mod)
		return b1 * b1 % mod
	} else {
		return powMod(b, p-1, mod) * b % mod
	}
}

func shortestWay(source string, target string) int {
	nt := len(target)
	j := 0
	cnt := 0
	for j < nt {
		s := longestSeq(source, target[j:])
		if s == 0 {
			return -1
		}
		cnt++
		j += s
	}
	return cnt
}

func longestSeq(source, target string) int {
	ns := len(source)
	nt := len(target)
	i := 0
	j := 0
	for i < ns && j < nt {
		for i < ns && source[i] != target[j] {
			i++
		}
		if i < ns {
			i++
			j++
		}
	}
	return j
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
