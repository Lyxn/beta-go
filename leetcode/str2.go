package leetcode

import (
	"math"
	"sort"
	"strconv"
)

func minimumDeleteSum(s1 string, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)
	ss1 := sumAscii(s1)
	ss2 := sumAscii(s2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if s1[i-1] == s2[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+int(s1[i-1]))
			}
		}
	}
	return ss1 + ss2 - 2*dp[n1][n2]
}

func sumAscii(ss string) int {
	res := 0
	for _, c := range ss {
		res += int(c)
	}
	return res
}

func splitIntoFibonacci(s string) []int {
	n := len(s)
	if n < 3 {
		return nil
	}
	for i := 1; i < n-1; i++ {
		f1, e := strconv.Atoi(s[0:i])
		if e != nil {
			break
		} else if s[0] == '0' && i > 1 {
			break
		}
		for j := i + 1; j < n; j++ {
			f2, e := strconv.Atoi(s[i:j])
			if e != nil || math.MaxInt32-f1 < f2 {
				break
			} else if s[i] == '0' && i+1 < j {
				break
			}
			res, ok := isValidFibo(s, i, j)
			if ok {
				return res
			}
		}
	}
	return nil
}

func isValidFibo(s string, a, b int) (res []int, ok bool) {
	n := len(s)
	f1, _ := strconv.Atoi(s[0:a])
	f2, _ := strconv.Atoi(s[a:b])
	res = append(res, f1, f2)
	j := b
	for j < n {
		if math.MaxInt32-f1 < f2 {
			return nil, false
		}
		f3 := f1 + f2
		s3 := strconv.Itoa(f3)
		for i := 0; i < len(s3); i++ {
			if j+i >= n || s3[i] != s[j+i] {
				return nil, false
			}
		}
		res = append(res, f3)
		f1, f2 = f2, f3
		j += len(s3)
	}
	return res, true
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	n3 := len(s3)
	if n1+n2 != n3 {
		return false
	}
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		if s3[i-1] == s1[i-1] {
			dp[i][0] = i
		} else {
			break
		}
	}
	for i := 1; i <= n2; i++ {
		if s3[i-1] == s2[i-1] {
			dp[0][i] = i
		} else {
			break
		}
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			n := dp[i-1][j]
			w := dp[i][j-1]
			if n > 0 && s3[n] == s1[i-1] {
				dp[i][j] = n + 1
			} else if w > 0 && s3[w] == s2[j-1] {
				dp[i][j] = w + 1
			}
		}
	}
	return dp[n1][n2] == n3
}

func numDistinct(s string, t string) int {
	ns := len(s)
	nt := len(t)
	if nt > ns {
		return 0
	} else if ns == 0 || nt == 0 {
		return 0
	}
	dp := make([]int, ns+1)
	dp1 := make([]int, ns+1)
	for j := 0; j < ns; j++ {
		if s[j] == t[0] {
			dp[j+1] += dp[j] + 1
		} else {
			dp[j+1] = dp[j]
		}
	}
	for i := 1; i < nt; i++ {
		for j := i; j < ns; j++ {
			if s[j] == t[i] {
				dp1[j+1] += dp1[j] + dp[j]
			} else {
				dp1[j+1] = dp1[j]
			}
		}
		dp, dp1 = dp1, dp
		resetInts(dp1)
		if dp[ns] == 0 {
			return 0
		}
	}
	return dp[ns]
}

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make(map[int]map[int]map[int]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]map[int]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make(map[int]bool, n)
		}
	}
	var recur func(l1, l2, ls int) bool
	recur = func(l1, l2, ls int) bool {
		if ls == 1 {
			return s1[l1] == s2[l2]
		}
		flag, ok := dp[l1][l2][ls]
		if ok {
			return flag
		}
		for i := 1; i < ls; i++ {
			flag = recur(l1, l2, i) && recur(l1+i, l2+i, ls-i)
			if flag {
				break
			}
			flag = recur(l1, l2+ls-i, i) && recur(l1+i, l2, ls-i)
			if flag {
				break
			}
		}
		dp[l1][l2][ls] = flag
		return flag
	}
	return recur(0, 0, n)
}

func isEqualBase(s1, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

func firstUniqChar(s string) int {
	cnt := make([]int, 26)
	pos := make([]int, 26)
	n := len(s)
	for i := 0; i < 26; i++ {
		pos[i] = n
	}
	for i := 0; i < n; i++ {
		idx := int(s[i] - 'a')
		cnt[idx] += 1
		if cnt[idx] > 1 {
			pos[idx] = n
		} else {
			pos[idx] = i
		}
	}
	mv := n
	for i := 0; i < 26; i++ {
		if pos[i] < mv {
			mv = pos[i]
		}
	}
	if mv == n {
		return -1
	}
	return mv
}
