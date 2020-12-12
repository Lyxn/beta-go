package leetcode

import "testing"

func TestMaxSizeSlices(t *testing.T) {
	tests := []struct {
		slices []int
		want   int
	}{
		{[]int{3, 1, 2}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 10},
		{[]int{8, 9, 8, 6, 1, 1}, 16},
		{[]int{4, 1, 2, 5, 8, 3, 1, 9, 7}, 21},
		{[]int{9, 8, 1, 7, 7, 9, 5, 10, 7, 9, 3, 8, 3, 4, 8}, 45},
		{[]int{1, 7, 7, 9, 5, 10, 7, 9, 3, 8, 3, 4}, 36},
		{[]int{7, 7, 9, 5, 10, 7, 9, 3, 8}, 28},
	}
	for _, tt := range tests {
		ret := maxSizeSlices(tt.slices)
		if ret != tt.want {
			t.Logf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestLargestDivisibleSubset(t *testing.T) {
	//nums := []int{1, 2, 3}
	//nums := []int{1, 2, 4, 5, 3, 8}
	//nums := []int{2,3,5,7,11,13,17,19,23,31,1000000007}
	nums := []int{4, 8, 10, 240}
	ret := largestDivisibleSubset(nums)
	t.Logf("get=%v", ret)
}

func TestSumOfDistancesInTree(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	n := len(edges) + 1
	ret := sumOfDistancesInTree(n, edges)
	t.Logf("get=%v", ret)
}

func TestSuperEggDrop(t *testing.T) {
	tests := []struct {
		k    int
		n    int
		want int
	}{
		{2, 6, 3},
		{3, 14, 4},
		{3, 25, 5},
	}
	for _, tt := range tests {
		ret := superEggDrop(tt.k, tt.n)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}

}
