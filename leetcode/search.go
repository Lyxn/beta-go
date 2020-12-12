package leetcode

import (
	"sort"
)

func solveNQueens(n int) [][]string {
	var ps [][]int
	for i := 0; i < n; i++ {
		ps = append(ps, []int{i})
	}
	for i := 1; i < n; i++ {
		var tps [][]int
		for j := 0; j < len(ps); j++ {
			rt := ps[j]
			for x := 0; x < n; x++ {
				if isValidQueen(rt, x) {
					rt1 := append([]int{}, rt...)
					rt1 = append(rt1, x)
					tps = append(tps, rt1)
				}
			}
		}
		ps = tps
	}
	return fmtQueens(ps)
}

func isValidQueen(rt []int, x int) bool {
	rx := len(rt)
	for r := 0; r < rx; r++ {
		c := rt[r]
		if x == c {
			return false
		} else if r-c == rx-x {
			return false
		} else if r+c == rx+x {
			return false
		}
	}
	return true
}

func fmtQueens(ps [][]int) [][]string {
	if len(ps) == 0 {
		return nil
	}
	n := len(ps[0])
	a := make([]rune, n)
	for i := 0; i < n; i++ {
		a[i] = '.'
	}
	var res [][]string
	for i := 0; i < len(ps); i++ {
		res = append(res, fmtQueen(ps[i]))
	}
	return res
}

func fmtQueen(rt []int) []string {
	n := len(rt)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		ss := make([]rune, n)
		for j := 0; j < n; j++ {
			if rt[i] == j {
				ss[j] = 'Q'
			} else {
				ss[j] = '.'
			}
		}
		res[i] = string(ss)
	}
	return res
}

func existWord(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && isExistAtXY(board, i, j, word) {
				return true
			}
		}
	}
	return false
}

func isExistAtXY(board [][]byte, x, y int, word string) bool {
	m := len(board)
	n := len(board[0])
	fb := make([][]bool, m)
	for i := 0; i < m; i++ {
		fb[i] = make([]bool, n)
	}
	nums := len(word)
	xys := [][]int{{x, y, 0, 0}}
	for len(xys) > 0 {
		xyz := xys[len(xys)-1]
		xys = xys[:len(xys)-1]
		i := xyz[0]
		j := xyz[1]
		k := xyz[2]
		z := xyz[3]
		if k == nums-1 {
			return true
		}
		if z == 1 {
			fb[i][j] = false
			continue
		}
		xys = append(xys, []int{i, j, k, 1})
		fb[i][j] = true
		k1 := k + 1
		if j-1 >= 0 && fb[i][j-1] != true && board[i][j-1] == word[k1] {
			xys = append(xys, []int{i, j - 1, k1, 0})
		}
		if j+1 < n && fb[i][j+1] != true && board[i][j+1] == word[k1] {
			xys = append(xys, []int{i, j + 1, k1, 0})
		}
		if i-1 >= 0 && fb[i-1][j] != true && board[i-1][j] == word[k1] {
			xys = append(xys, []int{i - 1, j, k1, 0})
		}
		if i+1 < m && fb[i+1][j] != true && board[i+1][j] == word[k1] {
			xys = append(xys, []int{i + 1, j, k1, 0})
		}
	}
	return false
}

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n <= 1 {
		return n
	}
	envelopes = append(envelopes, []int{0, 0})
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		} else {
			return envelopes[i][0] < envelopes[j][0]
		}
	})
	res := 0
	st := make([]int, n+1)
	st[0] = 0
	for i := 1; i <= n; i++ {
		xy := envelopes[i]
		mx := 0
		for j := 0; j < i; j++ {
			cur := envelopes[j]
			if cur[0] < xy[0] && cur[1] < xy[1] && mx < st[j] {
				mx = st[j]
			}
		}
		st[i] = mx + 1
		if st[i] > res {
			res = st[i]
		}
	}
	return res
}
