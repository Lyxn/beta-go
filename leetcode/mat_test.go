package leetcode

import "testing"

func TestGetMaxMatrix(t *testing.T) {
	tests := []struct {
		mat [][]int
	}{
		{
			[][]int{{0}, {1}, {1}},
		},
	}
	for _, tt := range tests {
		res := getMaxMatrix(tt.mat)
		t.Logf("sub=%v", res)
	}
}

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{2, 1, 5, 6, 2, 3},
			10,
		},
	}
	for _, tt := range tests {
		res := largestRectangleArea(tt.nums)
		if res != tt.want {
			t.Errorf("get=%v want=%v", res, tt.want)
		}
	}
}

func TestMatrixBlockSum(t *testing.T) {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 1
	ret := matrixBlockSum(mat, k)
	t.Logf("ret=%v", ret)
}

func TestMaxSideLength(t *testing.T) {
	mat := [][]int{{18, 70}, {61, 1}, {25, 85}, {14, 40}, {11, 96}, {97, 96}, {63, 45}}
	threshold := 40184
	ret := maxSideLength(mat, threshold)
	t.Logf("get=%v", ret)
}

func TestMaxSumSubmatrix(t *testing.T) {
	mat := [][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}
	k := 8
	ret := maxSumSubmatrix(mat, k)
	t.Logf("ret=%v", ret)
}

func TestMaximalSquare(t *testing.T) {
	mat := [][]byte{{1, 0, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}}
	ret := maximalSquare(mat)
	t.Logf("ret=%v", ret)
}

func TestNumSubmatrixSumTarget(t *testing.T) {
	mat := [][]int{{1, -1}, {-1, 1}}
	target := 0
	ret := numSubmatrixSumTarget(mat, target)
	t.Logf("ret=%v", ret)
}
