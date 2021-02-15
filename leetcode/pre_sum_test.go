package leetcode

import "testing"

func TestMinSumOfLengths(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{7, 3, 4, 7}, 7, 2},
		{[]int{4, 3, 2, 6, 2, 3, 4}, 6, -1},
		{[]int{3, 1, 1, 1, 5, 1, 2, 1}, 3, 3},
	}
	for _, tt := range tests {
		ret := minSumOfLengths(tt.nums, tt.k)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
func TestNumSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		//{[]int{3, 4}, 7, 2},
		//{[]int{10, 5, 2, 6}, 100, 8},
		//{[]int{10, 3, 3, 7, 2, 9, 7, 4, 7, 2, 8, 6, 5, 1, 5}, 30, 26},
		{[]int{5, 6}, 30, 2},
	}
	for _, tt := range tests {
		ret := numSubarrayProductLessThanK(tt.nums, tt.k)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
