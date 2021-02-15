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

func resetMat(dp [][]int, v int) {
	m := len(dp)
	n := len(dp[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = v
		}
	}
}

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp0 := make([][]int, m+1)
	dp1 := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp0[i] = make([]int, n+1)
		dp1[i] = make([]int, n+1)
	}
	maxVal := math.MaxInt32
	resetMat(dp0, maxVal)
	for i := 0; i <= n; i++ {
		dp0[0][i] = 0
	}
	for k := 1; k <= target; k++ {
		resetMat(dp1, maxVal)
		for i := k; i <= m; i++ {
			minC1 := 1
			minV1 := dp0[i-1][1]
			minV2 := maxVal
			for j := 2; j <= n; j++ {
				if minV1 > dp0[i-1][j] {
					minV2 = minV1
					minV1 = dp0[i-1][j]
					minC1 = j
				} else if minV2 > dp0[i-1][j] {
					minV2 = dp0[i-1][j]
				}
			}
			cc := houses[i-1]
			if cc != 0 {
				if cc == minC1 {
					dp1[i][cc] = min(dp1[i-1][cc], minV2)
				} else {
					dp1[i][cc] = min(dp1[i-1][cc], minV1)
				}
				continue
			}
			for j := 1; j <= n; j++ {
				mj := minV1
				if minC1 == j {
					mj = minV2
				}
				dp1[i][j] = min(dp1[i-1][j], mj)
				if dp1[i][j] < maxVal {
					dp1[i][j] += cost[i-1][j-1]
				}
			}
		}
		dp0, dp1 = dp1, dp0
	}
	minV := maxVal
	for j := 1; j <= n; j++ {
		if minV > dp0[m][j] {
			minV = dp0[m][j]
		}
	}
	if minV == maxVal {
		return -1
	}
	return minV
}

func ways(pizza []string, k int) (res int) {
	m := len(pizza)
	n := len(pizza[0])
	preSum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		preSum[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j]
			if pizza[i][j] == 'A' {
				preSum[i+1][j+1] += 1
			}
		}
	}
	if preSum[m][n] < k {
		return 0
	} else if k == 1 {
		return 1
	}
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k+1)
			for s := 0; s <= k; s++ {
				dp[i][j][s] = -1
			}
		}
	}
	mod := 1000000007
	var recur func(x, y, s int) int
	recur = func(x, y, s int) int {
		if dp[x][y][s] >= 0 {
			return dp[x][y][s]
		}
		res := preSum[m][n] + preSum[x][y] - preSum[x][n] - preSum[m][y]
		cnt := 0
		for i := x + 1; i < m; i++ {
			rx := preSum[i][n] + preSum[x][y] - preSum[x][n] - preSum[i][y]
			if rx > 0 && res-rx >= s-1 {
				if s > 2 {
					cnt += recur(i, y, s-1)
				} else {
					cnt += 1
				}
			}
		}
		for j := y + 1; j < n; j++ {
			ry := preSum[m][j] + preSum[x][y] - preSum[x][j] - preSum[m][y]
			if ry > 0 && res-ry >= s-1 {
				if s > 2 {
					cnt += recur(x, j, s-1)
				} else {
					cnt += 1
				}
			}
		}
		dp[x][y][s] = cnt % mod
		return dp[x][y][s]
	}
	return recur(0, 0, k)
}

func mergeStones(stones []int, K int) int {
	n := len(stones)
	if (n-1) % (K-1) != 0 {
		return -1
	}
	pre := make([]int, n+1)
	for i:=0; i<n; i++ {
		pre[i+1] = pre[i] + stones[i]
	}
	dp := make([][]int, n)
	for i:=0; i<n; i++ {
		dp[i] = make([]int, n)
	}
	for i:=n-2; i>=0;i-- {
		for j:=i+K-1; j<n; j++ {
			if j+1-i == K {
				dp[i][j] = pre[j+1] - pre[i]
				continue
			}
			dp[i][j] = math.MaxInt32
			for p:=i; p<j; p+=K-1 {
				dp[i][j] = min(dp[i][j], dp[i][p] + dp[p+1][j])
			}
			if (j-i) % (K-1) == 0 {
				dp[i][j] += pre[j+1] - pre[i]
			}
		}
	}
	return dp[0][n-1]
}