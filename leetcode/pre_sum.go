package leetcode

import "math"

func minSumOfLengths(arr []int, target int) int {
	preK := getPreK(arr, target)
	pstK := getPstK(arr, target)
	n := len(arr)
	mv := math.MaxInt32
	for i := 0; i < n; i++ {
		if preK[i+1] != math.MaxInt32 && pstK[i] != math.MaxInt32 {
			mv = min(mv, preK[i+1]+pstK[i+1])
		}
	}
	if mv == math.MaxInt32 {
		return -1
	}
	return mv
}

func getPreK(arr []int, target int) []int {
	n := len(arr)
	hs := make(map[int]int)
	hs[0] = -1
	pre := 0
	preK := make([]int, n+1)
	preK[0] = math.MaxInt32
	for i := 0; i < n; i++ {
		pre += arr[i]
		last, ok := hs[pre-target]
		if ok {
			preK[i+1] = min(preK[i], i-last)
		} else {
			preK[i+1] = preK[i]
		}
		hs[pre] = i
	}
	return preK
}

func getPstK(arr []int, target int) []int {
	n := len(arr)
	hs := make(map[int]int)
	hs[0] = n
	pst := 0
	pstK := make([]int, n+1)
	pstK[n] = math.MaxInt32
	for i := n - 1; i >= 0; i-- {
		pst += arr[i]
		last, ok := hs[pst-target]
		if ok {
			pstK[i] = min(pstK[i+1], last-i)
		} else {
			pstK[i] = pstK[i+1]
		}
		hs[pst] = i
	}
	return pstK
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	n := len(nums)
	dp := make([]float64, n+1)
	dp[0] = 0
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i] + math.Log(float64(nums[i]))
	}
	lnK := math.Log(float64(k)) - 1e-10
	cnt := 0
	for i := 0; i < n; i++ {
		l := i
		r := n
		for l < r {
			m := (l + r + 1) / 2
			if dp[m]-dp[i] < lnK {
				l = m
			} else {
				r = m - 1
			}
		}
		cnt += r - i
	}
	return cnt
}

func countTriplets(arr []int) int {
	n := len(arr)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i] ^ arr[i]
	}
	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if dp[j]^dp[i] == 0 {
				cnt += j - i - 1
			}
		}
	}
	return cnt
}
