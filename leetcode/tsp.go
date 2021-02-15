package leetcode

func getPathDist(A []string, pr [][]int, st []int) int {
	dst := len(A[st[0]])
	for i := 1; i < len(st); i++ {
		dst += len(A[st[i]]) - pr[st[i-1]][st[i]]
	}
	return dst
}

func buildOverlap(A []string) [][]int {
	n := len(A)
	pr := make([][]int, n)
	for i := 0; i < n; i++ {
		pr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			pr[i][j] = overlap(A[i], A[j])
		}
	}
	return pr
}

func overlap(s, t string) int {
	ns := len(s)
	nt := len(t)
	m := min(ns, nt)
	for ; m > 0; m-- {
		if s[ns-m:ns] == t[0:m] {
			return m
		}
	}
	return 0
}

func shortestSuperstringBack(A []string) string {
	n := len(A)
	pr := buildOverlap(A)

	minOrder := make([]int, n)
	minDist := len(A[0])
	for i := 1; i < n; i++ {
		minOrder[i] = i
		minDist += len(A[i]) - pr[i-1][i]
	}

	cs := append([]int{}, minOrder...)

	var back func(st []int, s int)
	back = func(st []int, s int) {
		if s == n {
			dst := getPathDist(A, pr, st)
			if dst < minDist {
				minDist = dst
				minOrder = append([]int{}, st...)
			}
			return
		}
		for i := s; i < n; i++ {
			st[s], st[i] = st[i], st[s]
			back(st, s+1)
			st[s], st[i] = st[i], st[s]
		}
	}
	back(cs, 0)

	res := A[minOrder[0]]
	for i := 1; i < n; i++ {
		s := A[minOrder[i]]
		ol := pr[minOrder[i-1]][minOrder[i]]
		res += s[ol:]
	}
	return res
}

func shortestSuperstring(A []string) string {
	n := len(A)
	ns := 1 << n
	ol := buildOverlap(A)
	dp := make([][]int, ns)
	pre := make([][]int, ns)
	for i := 0; i < ns; i++ {
		dp[i] = make([]int, n)
		pre[i] = make([]int, n)
		for j := 0; j < n; j++ {
			pre[i][j] = -1
		}
	}
	for s := 1; s < ns; s++ {
		for i := 0; i < n; i++ {
			if (s>>i)&1 == 0 {
				continue
			}
			lastSt := s ^ (1 << i)
			if lastSt == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if (lastSt>>j)&1 == 0 {
					continue
				}
				if pre[s][i] == -1 {
					dp[s][i] = dp[lastSt][j] + ol[j][i]
					pre[s][i] = j
				} else if dp[s][i] < dp[lastSt][j]+ol[j][i] {
					dp[s][i] = dp[lastSt][j] + ol[j][i]
					pre[s][i] = j
				}
			}
		}
	}
	mx := 0
	for i := 1; i < n; i++ {
		if dp[ns-1][i] >= dp[ns-1][mx] {
			mx = i
		}
	}
	ci := mx
	res := A[mx]
	st := ns - 1
	for st != 0 {
		pi := pre[st][ci]
		if pi == -1 {
			break
		}
		ov := ol[pi][ci]
		res = A[pi][:len(A[pi])-ov] + res
		st = st ^ (1 << ci)
		ci = pi
	}
	return res
}
