package leetcode

import "math"

func fib(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}
	base := [2][2]int{{1, 1}, {1, 0}}
	mat := powMat(base, N-1)
	return mat[0][0]
}

func mulMat(a, b [2][2]int) (ret [2][2]int) {
	ret[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	ret[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	ret[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	ret[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]
	return ret
}

func powMat(mat [2][2]int, n int) (ret [2][2]int) {
	if n == 0 {
		return [2][2]int{{1, 0}, {0, 1}}
	} else if n == 1 {
		return mat
	}
	if n%2 == 0 {
		sq := powMat(mat, n/2)
		return mulMat(sq, sq)
	} else {
		return mulMat(mat, powMat(mat, n-1))
	}
}

func getMaxMatrix(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m := len(matrix)
	n := len(matrix[0])
	r1, c1, r2, c2 := 0, 0, 0, 0
	max := matrix[0][0]
	for i := 0; i < m; i++ {
		arr := make([]int, n)
		for j := i; j < m; j++ {
			for k := 0; k < n; k++ {
				arr[k] += matrix[j][k]
			}
			sm, s1, s2 := subArray(arr)
			if sm > max {
				max = sm
				r1, r2, c1, c2 = i, j, s1, s2
			}
		}
	}
	return []int{r1, c1, r2, c2}
}

func subArray(nums []int) (max, c1, c2 int) {
	max = nums[0]
	c1 = 0
	c2 = 0
	lastC := 0
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		maxCur := nums[i]
		if last < 0 {
			lastC = i
		} else {
			maxCur += last
		}
		if maxCur > max {
			max = maxCur
			c1 = lastC
			c2 = i
		}
		last = maxCur
	}
	return max, c1, c2
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	sh := []int{0}
	ss := []int{0}
	max := 0
	for i := 1; i < len(heights); i++ {
		th := sh[len(sh)-1]
		if heights[i] > heights[th] {
			sh = append(sh, i)
			ss = append(ss, i)
			continue
		}
		cs := ss[len(ss)-1]
		for len(sh) > 0 && heights[i] <= heights[sh[len(sh)-1]] {
			h := sh[len(sh)-1]
			s := ss[len(ss)-1]
			area := heights[h] * (i - s)
			if area > max {
				max = area
			}
			cs = s
			sh = sh[:len(sh)-1]
			ss = ss[:len(ss)-1]
		}
		sh = append(sh, i)
		ss = append(ss, cs)
	}
	n := len(heights)
	for len(sh) > 0 {
		h := sh[len(sh)-1]
		s := ss[len(ss)-1]
		area := heights[h] * (n - s)
		if area > max {
			max = area
		}
		sh = sh[:len(sh)-1]
		ss = ss[:len(ss)-1]
	}
	return max
}

func getPreSum2(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	pre := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		pre[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			pre[i][j] += pre[i][j-1] + pre[i-1][j] + mat[i-1][j-1] - pre[i-1][j-1]
		}
	}
	return pre
}

func matrixBlockSum(mat [][]int, K int) [][]int {
	m := len(mat)
	n := len(mat[0])
	pre := getPreSum2(mat)
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = getBlock(pre, i, j, K)
		}
	}
	return res
}

func getBlock(pre [][]int, i, j, k int) int {
	m1 := len(pre)
	n1 := len(pre[0])
	n := i - k + 1
	s := i + k + 1
	w := j - k + 1
	e := j + k + 1
	if s >= m1 {
		s = m1 - 1
	}
	if n < 1 {
		n = 1
	}
	if w < 1 {
		w = 1
	}
	if e >= n1 {
		e = n1 - 1
	}
	res := 0
	res = pre[s][e] - pre[s][w-1] - pre[n-1][e] + pre[n-1][w-1]
	return res
}

func maxSideLength(mat [][]int, threshold int) int {
	if threshold == 0 {
		return 0
	}
	m := len(mat)
	n := len(mat[0])
	pre := getPreSum2(mat)
	l := 0
	r := min(m, n)
	for l < r {
		mid := (l + r + 1) / 2
		sq := getMinSquare(pre, mid)
		if sq <= threshold {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func getMinSquare(pre [][]int, k int) int {
	if k == 0 {
		return 0
	}
	m := len(pre)
	n := len(pre[0])
	mn := math.MaxInt32
	for i := k; i < m; i++ {
		for j := k; j < n; j++ {
			sq := pre[i][j] - pre[i-k][j] - pre[i][j-k] + pre[i-k][j-k]
			mn = min(sq, mn)
		}
	}
	return mn
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	nx := math.MinInt32
	for l := 0; l < n; l++ {
		row := make([]int, m)
		for r := l; r < n; r++ {
			for i := 0; i < m; i++ {
				row[i] += matrix[i][r]
			}
			a := findMax(row, k)
			if a == k {
				return k
			} else if a > nx {
				nx = a
			}
		}
	}
	return nx
}

func findMax(row []int, k int) int {
	nx := math.MinInt32
	dp := 0
	for i := 0; i < len(row); i++ {
		if dp > 0 {
			dp += row[i]
		} else {
			dp = row[i]
		}
		if dp == k {
			return k
		}
		v := dp
		if v > k {
			v = findMax1(row[:i+1], k)
		}
		if v == k {
			return v
		} else if v > nx {
			nx = v
		}
	}
	return nx
}

func findMax1(row []int, k int) int {
	s := 0
	nx := math.MinInt32
	for i := len(row) - 1; i > 0; i-- {
		s += row[i]
		if s > k {
			continue
		} else if s == k {
			return k
		} else if s > nx {
			nx = s
		}
	}
	return nx
}
