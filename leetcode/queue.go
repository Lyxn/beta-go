package leetcode

import (
	"bytes"
	"fmt"
)

type Pair struct {
	x int
	y int
}

func PrintMatrix(mat [][]int) {
	for _, v := range mat {
		fmt.Printf("%v\n", v)
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	total := m * n
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	offset := 0
	cnt := 0
	for offset < total {
		x := offset / n
		y := offset % n
		isFind := false
	Loop:
		for i := x; i < m; i++ {
			for j := 0; j < n; j++ {
				if vis[i][j] {
					continue
				}
				vis[i][j] = true
				if grid[i][j] == 1 {
					isFind = true
					x = i
					y = j
					break Loop
				}
			}
		}
		if !isFind {
			break
		}
		q := []Pair{
			{x, y},
		}
		for len(q) > 0 {
			sz := len(q)
			for i := 0; i < sz; i++ {
				r := q[i]
				q = fill(q, vis, grid, r)
			}
			q = q[sz:]
		}
		cnt++
		offset = x*n + y + 1
	}
	return cnt
}

func fill(q []Pair, vis [][]bool, grid [][]byte, r Pair) []Pair {
	m := len(grid)
	n := len(grid[0])
	prList := []Pair{
		{r.x, r.y - 1},
		{r.x, r.y + 1},
		{r.x - 1, r.y},
		{r.x + 1, r.y},
	}
	for _, p := range prList {
		if p.x < 0 || p.y < 0 || p.x >= m || p.y >= n {
			continue
		} else if vis[p.x][p.y] {
			continue
		}
		vis[p.x][p.y] = true
		if grid[p.x][p.y] == 1 {
			q = append(q, p)
		}
	}
	return q
}

func sqrtInt(n int) (x int) {
	x = 1
	for i := 0; i < 1e3; i++ {
		x2 := x * x
		if x2 <= n && x2+2*x+1 > n {
			break
		} else {
			x = (n + x2) / x / 2
		}
	}
	return x
}

func openLock(deadends []string, target string) int {
	s0 := "0000"
	gd := make(map[string]struct{}, len(deadends))
	for _, k := range deadends {
		gd[k] = struct{}{}
	}
	if _, ok := gd[target]; ok {
		return -1
	} else if _, ok0 := gd[s0]; ok0 {
		return -1
	}
	if target == s0 {
		return 0
	}
	gd[s0] = struct{}{}
	qn := []string{s0}
	step := 1
	for len(qn) > 0 {
		sz := len(qn)
		//fmt.Printf("qn=%v gd=%v\n", len(qn), len(gd))
		for i := 0; i < sz; i++ {
			k := qn[i]
			for _, s := range chg(k) {
				if s == target {
					return step
				}
				_, okd := gd[s]
				if okd {
					continue
				} else {
					gd[s] = struct{}{}
				}
				qn = append(qn, s)
			}
		}
		qn = qn[sz:]
		step++
	}
	return -1
}

func chg(s string) (strs []string) {
	n := len(s)
	buf := bytes.Buffer{}
	for i := 0; i < n; i++ {
		var p, n rune
		if s[i] == '0' {
			p = '9'
			n = '1'
		} else if s[i] == '9' {
			p = '8'
			n = '0'
		} else {
			p = rune(s[i] - 1)
			n = rune(s[i] + 1)
		}
		buf.Reset()
		buf.WriteString(s[:i])
		buf.WriteRune(p)
		buf.WriteString(s[i+1:])
		sp := buf.String()
		strs = append(strs, sp)
		buf.Reset()
		buf.WriteString(s[:i])
		buf.WriteRune(n)
		buf.WriteString(s[i+1:])
		sn := buf.String()
		strs = append(strs, sn)
	}
	return strs
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	nr := len(image)
	nc := len(image[0])
	base := image[sr][sc]
	if base == newColor {
		return image
	}
	image[sr][sc] = newColor
	qr := []int{sr}
	qc := []int{sc}
	for len(qr) > 0 {
		r := qr[0]
		c := qc[0]
		qr = qr[1:]
		qc = qc[1:]
		if r-1 >= 0 && image[r-1][c] == base {
			image[r-1][c] = newColor
			qr = append(qr, r-1)
			qc = append(qc, c)
		}
		if r+1 < nr && image[r+1][c] == base {
			image[r+1][c] = newColor
			qr = append(qr, r+1)
			qc = append(qc, c)
		}
		if c-1 >= 0 && image[r][c-1] == base {
			image[r][c-1] = newColor
			qr = append(qr, r)
			qc = append(qc, c-1)
		}
		if c+1 < nc && image[r][c+1] == base {
			image[r][c+1] = newColor
			qr = append(qr, r)
			qc = append(qc, c+1)
		}
	}
	return image
}

func updateMatrix(matrix [][]int) [][]int {
	nr := len(matrix)
	nc := len(matrix[0])
	dist := make([][]int, nr)
	seen := make([][]bool, nr)
	for i := 0; i < nr; i++ {
		dist[i] = make([]int, nc)
		seen[i] = make([]bool, nc)
	}
	var qr []int
	var qc []int
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if matrix[i][j] == 0 {
				seen[i][j] = true
			} else if isEdge(matrix, i, j) && !seen[i][j] {
				seen[i][j] = true
				dist[i][j] = 1
				qr = append(qr, i)
				qc = append(qc, j)
			}
		}
	}
	for len(qr) > 0 {
		i := qr[0]
		j := qc[0]
		qr = qr[1:]
		qc = qc[1:]
		step := dist[i][j] + 1
		if i-1 >= 0 && !seen[i-1][j] {
			seen[i-1][j] = true
			dist[i-1][j] = step
			qr = append(qr, i-1)
			qc = append(qc, j)
		}
		if i+1 < nr && !seen[i+1][j] {
			seen[i+1][j] = true
			dist[i+1][j] = step
			qr = append(qr, i+1)
			qc = append(qc, j)
		}
		if j-1 >= 0 && !seen[i][j-1] {
			seen[i][j-1] = true
			dist[i][j-1] = step
			qr = append(qr, i)
			qc = append(qc, j-1)
		}
		if j+1 < nc && !seen[i][j+1] {
			seen[i][j+1] = true
			dist[i][j+1] = step
			qr = append(qr, i)
			qc = append(qc, j+1)
		}
	}
	return dist
}

func isEdge(mat [][]int, i, j int) bool {
	nr := len(mat)
	nc := len(mat[0])
	if mat[i][j] == 0 {
		return false
	}
	if i-1 >= 0 && mat[i-1][j] == 0 {
		return true
	} else if i+1 < nr && mat[i+1][j] == 0 {
		return true
	} else if j-1 >= 0 && mat[i][j-1] == 0 {
		return true
	} else if j+1 < nc && mat[i][j+1] == 0 {
		return true
	}
	return false
}
