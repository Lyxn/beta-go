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
