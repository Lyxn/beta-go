package leetcode

import (
	"sort"
)

func maxSizeSlices(slices []int) int {
	n := len(slices)
	cs := n / 3
	return max(dpMaxSizeSlices(slices[:n-1], cs), dpMaxSizeSlices(slices[1:], cs))
}

func dpMaxSizeSlices(nums []int, cs int) int {
	n := len(nums)
	dp0 := make([]int, n+1)
	for j := 1; j <= n; j++ {
		dp0[j] = max(nums[j-1], dp0[j-1])
	}
	for i := 2; i <= cs; i++ {
		dp1 := make([]int, n+1)
		dp1[1] = nums[0]
		for j := 2; j <= n; j++ {
			dp1[j] = max(dp0[j-2]+nums[j-1], dp1[j-1])
		}
		dp0 = dp1
	}
	return dp0[n]
}

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	sort.Ints(nums)
	dpCnt := make([]int, n)
	dpIdx := make([]int, n)
	for i := 0; i < n; i++ {
		dpIdx[i] = -1
	}
	for i := 0; i < n-1; i++ {
		cnt := dpCnt[i] + 1
		s := i + 1
		k := nums[i]
		for s < n {
			if nums[s]%nums[i] == 0 {
				if dpCnt[s] < cnt {
					dpCnt[s] = cnt
					dpIdx[s] = i
				}
				s++
				continue
			}
			k = (nums[s]/nums[i] + 1) * nums[i]
			if k > nums[n-1] {
				break
			}
			j := search2(nums, s, k)
			if j < n && nums[j] == k {
				if dpCnt[j] < cnt {
					dpCnt[j] = cnt
					dpIdx[j] = i
				}
				j++
			}
			s = j
		}
	}
	mi := 0
	mc := 0
	for i := 1; i < n; i++ {
		if dpCnt[i] > mc {
			mc = dpCnt[i]
			mi = i
		}
	}
	rs := []int{}
	for mi >= 0 {
		rs = append(rs, nums[mi])
		mi = dpIdx[mi]
	}
	return rs
}

func search2(nums []int, s, k int) int {
	l := s
	r := len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func sumOfDistancesInTree0(N int, edges [][]int) []int {
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}
	nxs := make([][]int, N)
	for _, edge := range edges {
		s := edge[0]
		d := edge[1]
		dp[s][d] = 1
		dp[d][s] = 1
		nxs[s] = append(nxs[s], d)
		nxs[d] = append(nxs[d], s)
	}

	fillShortestDist(dp)

	res := make([]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			res[i] += dp[i][j]
		}
	}
	return res
}

func fillShortestDist(dp [][]int) {
	N := len(dp)
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				if dp[i][j] > 0 {
					continue
				} else if i == k || j == k {
					continue
				}
				if dp[i][k] > 0 && dp[k][j] > 0 {
					dp[i][j] = dp[i][k] + dp[k][j]
					dp[j][i] = dp[i][k] + dp[k][j]
				}
			}
		}
	}
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	sz := make([]int, n)
	dp := make([]int, n)
	var dfs func(u, f int)
	dfs = func(u, f int) {
		sz[u] = 1
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			dfs(v, u)
			dp[u] += dp[v] + sz[v]
			sz[u] += sz[v]
		}
	}
	dfs(0, -1)

	ans := make([]int, n)
	var dfs2 func(u, f int)
	dfs2 = func(u, f int) {
		ans[u] = dp[u]
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			pu, pv := dp[u], dp[v]
			su, sv := sz[u], sz[v]

			dp[u] -= dp[v] + sz[v]
			sz[u] -= sz[v]
			dp[v] += dp[u] + sz[u]
			sz[v] += sz[u]

			dfs2(v, u)

			dp[u], dp[v] = pu, pv
			sz[u], sz[v] = su, sv
		}
	}
	dfs2(0, -1)
	return ans
}

func superEggDrop(K int, N int) int {
	dp := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dp[i] = i
	}
	for k := 2; k <= K; k++ {
		dk := make([]int, N+1)
		dk[1] = 1
		for i := 2; i <= N; i++ {
			dk[i] = searchMinMax(dp[:i], dk[:i]) + 1
		}
		dp = dk
	}
	return dp[N]
}

func searchMinMax(dp, dk []int) int {
	n := len(dp)
	if n <= 2 {
		return max(dp[n-1], dk[n-1])
	}
	l := 0
	r := n - 2
	for l <= r {
		m := (l + r + 1) / 2
		nm := n - 1 - m
		if dp[m] < dk[nm] && dp[m+1] >= dk[nm-1] {
			return min(dp[m+1], dk[nm])
		}
		if dp[m] < dk[nm] {
			l = m
		} else {
			r = m - 1
		}
	}
	return -1
}
