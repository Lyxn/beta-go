package leetcode

import "sort"

func minCostConnectPointsUF(points [][]int) int {
	n := len(points)

	md := func(i, j int) int {
		return abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])
	}

	arcs := make([][3]int, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			arcs = append(arcs, [3]int{i, j, md(i, j)})
		}
	}
	sort.Slice(arcs, func(i, j int) bool {
		return arcs[i][2] < arcs[j][2]
	})

	pre := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = i
		cnt[i] = 1
	}

	find := func(i int) int {
		c := i
		for pre[c] != c {
			c = pre[c]
		}
		for pre[i] != c {
			i = pre[i]
			pre[i] = c
		}
		return c
	}

	union := func(i, j int) (int, bool) {
		pi := find(i)
		pj := find(j)
		if pi != pj {
			pre[pi] = pj
			cnt[pj] += cnt[pi]
		}
		return cnt[pj], pi != pj
	}

	res := 0
	x := 1
	c := 0
	ok := false
	for x < n && c < len(arcs) {
		arc := arcs[c]
		x, ok = union(arc[0], arc[1])
		if ok {
			res += arc[2]
		}
		c++
	}
	return res
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)

	md := func(i, j int) int {
		return abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])
	}

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist[i][j] = md(i, j)
			dist[j][i] = dist[i][j]
		}
	}

	costs := make([][3]int, 0)
	treeNodes := make([]bool, n)
	treeNodes[0] = true
	for i := 1; i < n; i++ {
		costs = append(costs, [3]int{0, i, dist[0][i]})
	}
	res := 0
	for j := 1; j < n; j++ {
		mi := 0
		for i := 1; i < len(costs); i++ {
			if costs[i][2] < costs[mi][2] {
				mi = i
			}
		}
		mc := costs[mi]
		res += mc[2]
		nv := mc[1]
		treeNodes[nv] = true
		l := 0
		v := 0
		for ; l < len(costs); l++ {
			if costs[l][1] == nv {
				continue
			}
			costs[v] = costs[l]
			v++
		}
		costs = costs[:v]
		for i := 1; i < n; i++ {
			if treeNodes[i] {
				continue
			}
			costs = append(costs, [3]int{nv, i, dist[nv][i]})
		}
	}
	return res
}
