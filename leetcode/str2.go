package leetcode

import (
	"math"
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
	for i:=1; i<n-1; i++ {
		f1, e := strconv.Atoi(s[0:i])
		if e != nil {
			break
		} else if s[0] == '0' && i > 1 {
			break
		}
		for j:=i+1; j<n; j++ {
			f2, e := strconv.Atoi(s[i:j])
			if e != nil || math.MaxInt32 - f1 < f2 {
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
		if math.MaxInt32 - f1 < f2 {
			return nil, false
		}
		f3 := f1 + f2
		s3 := strconv.Itoa(f3)
		for i:=0; i<len(s3); i++ {
			if j+i>=n || s3[i] != s[j+i] {
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
	if n1 + n2 != n3 {
		return false
	}
	dp := make([][]int, n1+1)
	for i:=0; i<=n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	for i:=1; i<=n1; i++ {
		if s3[i-1] == s1[i-1] {
			dp[i][0] = i
		} else {
			break
		}
	}
	for i:=1; i<=n2; i++ {
		if s3[i-1] == s2[i-1] {
			dp[0][i] = i
		} else {
			break
		}
	}

	for i:=1; i<=n1; i++ {
		for j:=1; j<=n2; j++ {
			n := dp[i-1][j]
			w := dp[i][j-1]
			if n > 0 && s3[n] == s1[i-1] {
				dp[i][j] = n+1
			} else if w > 0 && s3[w] == s2[j-1] {
				dp[i][j] = w+1
			}
		}
	}
	return dp[n1][n2] == n3
}
