package leetcode

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, s := range stones {
		sum += s
	}
	target := sum / 2
	max := 0
	dp := make([]bool, target+1)
	dp[0] = true
	for _, s := range stones {
		for j := target; j >= s; j-- {
			dp[j] = dp[j-s] || dp[j]
		}
	}
	for j := target; j >= 0; j-- {
		if dp[j] == true {
			max = j
			break
		}
	}
	return sum - max*2
}

func profitableSchemes(G int, P int, group []int, profit []int) int {
	n := len(group)
	if n == 0 {
		return 0
	}
	sp := 0
	for i := 0; i < n; i++ {
		sp += profit[i]
	}
	if sp < P {
		return 0
	}
	dp := make([][]int, G+1)
	for i := 0; i <= G; i++ {
		dp[i] = make([]int, sp+1)
	}
	dp[0][0] = 1
	md := 1000000007
	mx := 0
	for k := 0; k < n; k++ {
		gk := group[k]
		pk := profit[k]
		mx += pk
		for i := G; i >= gk; i-- {
			for j := mx; j >= pk; j-- {
				dp[i][j] += dp[i-gk][j-pk]
				if dp[i][j] > md {
					dp[i][j] -= md
				}
			}
		}
	}
	cnt := 0
	for j := sp; j >= P; j-- {
		for i := 1; i <= G; i++ {
			cnt += dp[i][j]
			if cnt > md {
				cnt -= md
			}
		}
	}
	return cnt
}

func profitableSchemes1(G int, P int, group []int, profit []int) int {
	n := len(group)
	if n == 0 {
		return 0
	}
	sp := 0
	for i := 0; i < n; i++ {
		sp += profit[i]
	}
	if sp < P {
		return 0
	}
	dp := make([][]int, G+1)
	dp1 := make([][]int, G+1)
	for i := 0; i <= G; i++ {
		dp[i] = make([]int, P+1)
		dp1[i] = make([]int, P+1)
	}
	dp[0][0] = 1
	md := 1000000007
	for k := 0; k < n; k++ {
		gk := group[k]
		pk := profit[k]
		copyMat(dp, dp1)
		for i := gk; i <= G; i++ {
			for j := 0; j <= P; j++ {
				j1 := j + pk
				if j1 > P {
					j1 = P
				}
				dp1[i][j1] += dp[i-gk][j]
				if dp1[i][j1] > md {
					dp1[i][j1] -= md
				}
			}
		}
		dp, dp1 = dp1, dp
	}
	cnt := 0
	for i := 1; i <= G; i++ {
		cnt += dp[i][P]
		if cnt > md {
			cnt -= md
		}
	}
	return cnt
}

func copyMat(src, dst [][]int) {
	m := len(src)
	n := len(src[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dst[i][j] = src[i][j]
		}
	}
}

func wordsTyping(sentence []string, rows int, cols int) int {
	nw := len(sentence)
	dt := make([]int, nw)
	dc := make([]int, nw)
	for i := 0; i < nw; i++ {
		dt[i] = -1
	}
	s := 0
	for dt[s] < 0 {
		si := s
		rc := cols
		cnt := 0
		for rc > 0 && len(sentence[si]) <= rc {
			if si == nw-1 {
				cnt++
			}
			rc -= len(sentence[si]) + 1
			si++
			if si == nw {
				si = 0
			}
		}
		dt[s] = si
		dc[s] = cnt
		s = si
	}
	cycleS := s
	cntPre := 0
	s = 0
	startR := 0
	for s != cycleS && startR < rows {
		cntPre += dc[s]
		s = dt[s]
		startR++
	}
	if startR == rows {
		return cntPre
	}
	cycle := 1
	cycleCnt := dc[cycleS]
	s = dt[cycleS]
	for s != cycleS {
		cycle++
		cycleCnt += dc[s]
		s = dt[s]
	}
	cntR := (rows - startR) / cycle
	cnt := cntPre + cycleCnt*cntR
	rr := startR + cycle*cntR
	s = cycleS
	for r := rr; r < rows; r++ {
		cnt += dc[s]
		s = dt[s]
	}
	return cnt
}

func wordsTyping1(sentence []string, rows int, cols int) int {
	mx := 0
	for _, l := range sentence {
		if len(l) > mx {
			mx = len(l)
		}
	}
	if mx > cols {
		return 0
	}
	nw := len(sentence)
	idx0 := make([]int, nw)
	for i := 0; i < nw; i++ {
		idx0[i] = -1
	}
	r := 0
	si := 0
	rc := cols
	cycle := -1
	cycleIdx := -1
	cycleCnt := 0
	cnt := 0
	for ; r < rows; r++ {
		if idx0[si] < 0 {
			idx0[si] = r
		} else {
			cycle = r - idx0[si]
			cycleIdx = si
			break
		}
		rc = cols
		for rc > 0 && len(sentence[si]) <= rc {
			if si == nw-1 {
				cnt++
			}
			rc -= len(sentence[si]) + 1
			si++
			if si == nw {
				si = 0
			}
		}
	}

	if cycle < 0 {
		return cnt
	}
	cntC := (rows - 1 - idx0[cycleIdx]) / cycle
	rr := idx0[cycleIdx] + cntC*cycle
	rCnt := 0

	si = cycleIdx
	for r = 0; r < rows; r++ {
		if r != 0 && si == cycleIdx {
			break
		}
		rc = cols
		for rc > 0 && len(sentence[si]) <= rc {
			if si == nw-1 {
				cycleCnt++
			}
			rc -= len(sentence[si]) + 1
			si++
			if si == nw {
				si = 0
			}
		}
		if rr+r == rows-1 {
			rCnt = cycleCnt
		}
	}
	cntTotal := cnt - cycleCnt + cycleCnt*cntC + rCnt
	return cntTotal
}
