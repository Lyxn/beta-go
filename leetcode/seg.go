package leetcode

import "sort"

type Seg struct {
	Val   int
	Li    int
	Ri    int
	Left  *Seg
	Right *Seg
}

func (s *Seg) IsLeaf() bool {
	return s.Left == nil && s.Right == nil
}

func (s *Seg) SetLeaf() {
	s.Left = nil
	s.Right = nil
}

func (s *Seg) GetMid() int {
	return (s.Li + s.Ri) / 2
}

func AddSeg(root *Seg, li, ri, height int) *Seg {
	if li >= ri {
		return root
	}
	if root.Ri <= li {
		return root
	} else if root.Li >= ri {
		return root
	}
	if root.IsLeaf() && li <= root.Li && root.Ri <= ri {
		if root.Val < height {
			root.Val = height
		}
		return root
	} else if root.IsLeaf() && root.Val >= height {
		return root
	}
	mid := root.GetMid()
	if !(mid <= li) {
		if root.Left == nil {
			root.Left = &Seg{Li: root.Li, Ri: mid, Val: root.Val}
		}
		root.Left = AddSeg(root.Left, li, ri, height)
	}
	if !(mid >= ri) {
		if root.Right == nil {
			root.Right = &Seg{Li: mid, Ri: root.Ri, Val: root.Val}
		}
		root.Right = AddSeg(root.Right, li, ri, height)
	}
	return root
}

func dfsSeg(root *Seg) [][]int {
	if root == nil {
		return nil
	} else if root.IsLeaf() {
		return [][]int{{root.Li, root.Val}, {root.Ri, 0}}
	}
	var lhs, rhs [][]int
	if root.Left == nil {
		if root.Val != 0 {
			lhs = [][]int{{root.Li, root.Val}, {root.GetMid(), 0}}
		}
	} else {
		lhs = dfsSeg(root.Left)
	}
	if root.Right == nil {
		if root.Val != 0 {
			rhs = [][]int{{root.GetMid(), root.Val}, {root.Ri, 0}}
		}
	} else {
		rhs = dfsSeg(root.Right)
	}
	if len(lhs) == 0 {
		return rhs
	} else if len(rhs) == 0 {
		return lhs
	}
	res := make([][]int, 0, len(lhs)+len(rhs))
	nl := len(lhs)
	if lhs[nl-1][0] == rhs[0][0] {
		lhs = lhs[:nl-1]
		if lhs[nl-2][1] == rhs[0][1] {
			rhs = rhs[1:]
		}
	}
	res = append(res, lhs...)
	res = append(res, rhs...)
	return res
}

/*
Param:
	buildings []Struct{Li, Ri, Height}
*/
func getSkyline(buildings [][]int) (res [][]int) {
	if len(buildings) == 0 {
		return nil
	}
	idxs := sortUniq(buildings)
	root := &Seg{Li: idxs[0], Ri: idxs[len(idxs)-1]}
	for i := 0; i < len(buildings); i++ {
		b := buildings[i]
		root = AddSeg(root, b[0], b[1], b[2])
	}
	res = dfsSeg(root)
	return res
}

func sortUniq(buildings [][]int) []int {
	n := len(buildings)
	dct := make(map[int]struct{}, n)
	for _, b := range buildings {
		dct[b[0]] = struct{}{}
		dct[b[1]] = struct{}{}
	}
	res := make([]int, 0, len(dct))
	for k := range dct {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	a0 := findInterval(intervals, newInterval[0])
	a1 := findInterval(intervals, newInterval[1])
	if a0 == n {
		return append(intervals, newInterval)
	}
	var res [][]int
	res = append(res, intervals[:a0]...)
	if a0 == a1 && intervals[a0][0] > newInterval[0] {
		res = append(res, newInterval)
		res = append(res, intervals[a0:]...)
		return res
	}
	lo := min(intervals[a0][0], newInterval[0])
	if a1 == n {
		hi := newInterval[1]
		res = append(res, []int{lo, hi})
	} else if newInterval[1] < intervals[a1][0] {
		hi := newInterval[1]
		res = append(res, []int{lo, hi})
		res = append(res, intervals[a1:]...)
	} else {
		hi := max(intervals[a1][1], newInterval[1])
		res = append(res, []int{lo, hi})
		if a1+1 < n {
			res = append(res, intervals[a1+1:]...)
		}
	}
	return res
}

func findInterval(intervals [][]int, x int) int {
	n := len(intervals)
	l := 0
	r := n
	for l < r {
		m := (l + r) / 2
		if intervals[m][1] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
