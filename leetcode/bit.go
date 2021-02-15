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

func countTripletsBit(A []int) int {
	na := len(A)
	ns := 1<<16
	dp := make([]int, ns)
	for i := 0; i < na; i++ {
		for j := 0; j < na; j++ {
			s := A[i] & A[j]
			dp[s]++
		}
	}
	res := 0
	for i := 0; i < na; i++ {
		st := ns - 1 - A[i]
		res += dp[0]
		for s := st; s != 0; s = (s - 1) & st {
			res += dp[s]
		}
	}
	return res
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	//ns := 1 << 20
	if desiredTotal == 0 {
		return true
	}
	if desiredTotal > maxChoosableInteger * (maxChoosableInteger + 1) / 2 {
		return false
	}
	dp := make([]map[int]bool, desiredTotal+1)
	for i:= 0; i<=desiredTotal; i++ {
		dp[i] = make(map[int]bool)
	}
	var dfs func(stt, ttl int) bool
	dfs = func(stt, ttl int) bool {
		if ttl <= 0 {
			return false
		}
		res, ok := dp[ttl][stt]
		if ok {
			return res
		}
		canWin := false
		for i := maxChoosableInteger-1; i >= 0; i-- {
			if (stt >> i) & 1 > 0 {
				continue
			}
			nxt := stt | (1 << i)
			win := dfs(nxt, ttl - i - 1)
			if win == false {
				canWin = true
				break
			}
		}
		dp[ttl][stt] = canWin
		return canWin
	}
	return dfs(0, desiredTotal)
}