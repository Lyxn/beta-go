package leetcode

import (
	"math"
	"sort"
)

func minSwap(A []int, B []int) int {
	n := len(A)
	dp0 := make([]int, n)
	dp1 := make([]int, n)
	for i := 1; i < n; i++ {
		dp0[i] = math.MaxInt32
		dp1[i] = math.MaxInt32
	}
	dp1[0] = 1
	for i := 1; i < n; i++ {
		if A[i] > A[i-1] && B[i] > B[i-1] {
			dp0[i] = dp0[i-1]
			dp1[i] = dp1[i-1] + 1
		}
		if A[i] > B[i-1] && B[i] > A[i-1] {
			dp0[i] = min(dp0[i], dp1[i-1])
			dp1[i] = min(dp1[i], dp0[i-1]+1)
		}
	}
	return min(dp0[n-1], dp1[n-1])
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	vis := make([][]int, 0)
	ns := len(stations)
	stop := startFuel
	cnt := 0
	nxt := 0
	for stop < target {
		for ; nxt < ns; nxt++ {
			if stations[nxt][0] > stop {
				break
			}
			vis = append(vis, stations[nxt])
			bottomUpVis(vis)
		}
		nv := len(vis)
		if nv == 0 {
			cnt = -1
			break
		}
		stop += vis[0][1]
		vis[0] = vis[nv-1]
		vis = vis[:nv-1]
		topDownVis(vis, 0)
		cnt++
	}
	return cnt
}

func bottomUpVis(vis [][]int) {
	c := len(vis) - 1
	for c > 0 {
		p := (c - 1) / 2
		if vis[p][1] >= vis[c][1] {
			break
		}
		vis[p], vis[c] = vis[c], vis[p]
		c = p
	}
}

func topDownVis(vis [][]int, p int) {
	n := len(vis)
	l := 2*p + 1
	r := 2*p + 2
	for l < n {
		mx := p
		if vis[mx][1] < vis[l][1] {
			mx = l
		}
		if r < n && vis[mx][1] < vis[r][1] {
			mx = r
		}
		if mx == p {
			break
		}
		vis[mx], vis[p] = vis[p], vis[mx]
		p = mx
		l = 2*p + 1
		r = 2*p + 2
	}
}

func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)
	dp := make([]float64, n+1)
	pre := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + float64(A[i-1])
		dp[i] = pre[i] / float64(i)
	}
	for k := 2; k <= K; k++ {
		dp1 := make([]float64, n+1)
		dp1[k] = pre[k]
		for i := k + 1; i <= n; i++ {
			mx := dp[k-1]
			for j := k - 1; j < i; j++ {
				avg := (pre[i] - pre[j]) / float64(i-j)
				mx = maxFloat(mx, dp[j]+avg)
			}
			dp1[i] = mx
		}
		dp = dp1
	}
	return dp[n]
}

func probabilityOfHeads(prob []float64, target int) float64 {
	n := len(prob)
	dp := make([]float64, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] * (1 - prob[i-1])
	}
	for k := 1; k <= target; k++ {
		dp1 := make([]float64, n+1)
		dp1[k] = dp[k-1] * prob[k-1]
		for i := k + 1; i <= n; i++ {
			dp1[i] = dp1[i-1]*(1-prob[i-1]) + dp[i-1]*prob[i-1]
		}
		dp = dp1
	}
	return dp[n]
}

func minDistance(houses []int, k int) int {
	sort.Ints(houses)
	n := len(houses)
	d1 := make([][]int, n)
	for i := 0; i < n; i++ {
		d1[i] = make([]int, n)
		for j := i; j < n; j++ {
			d1[i][j] = minDist1(houses[i : j+1])
		}
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = d1[0][i]
	}
	for j := 1; j < k; j++ {
		dp1 := make([]int, n)
		dp1[j] = 0
		for r := j + 1; r < n; r++ {
			minDist := math.MaxInt32
			for l := j - 1; l < r; l++ {
				minDist = min(minDist, dp[l]+d1[l+1][r])
			}
			dp1[r] = minDist
		}
		dp = dp1
	}
	return dp[n-1]
}

func minDist1(houses []int) int {
	n := len(houses)
	dist := 0
	l := 0
	r := n - 1
	for l < r {
		dist += houses[r] - houses[l]
		l++
		r--
	}
	return dist
}
