package leetcode

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{4, 10, 4, 3, 8, 9}, 3},
	}
	for _, tt := range tests {
		ret := lengthOfLIS(tt.nums)
		if ret != tt.want {
			t.Logf("wamt=%v get=%v", tt.want, ret)
		}
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 2}, []int{-1, 3}, 1.5},
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{3}, []int{-2, -1}, -1},
	}
	for _, tt := range tests {
		res := findMedianSortedArrays(tt.nums1, tt.nums2)
		if tt.want != res {
			t.Errorf("want=%v get=%v", tt.want, res)
		}
	}

}

func TestBisect1(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2}, 0, 0},
		{[]int{1, 2}, 2, 2},
		{[]int{1, 2}, 1, 1},
		{[]int{1, 2, 4, 9, 21}, 4, 3},
	}
	for _, tt := range tests {
		res := bisect1(tt.nums, tt.k)
		if tt.want != res {
			t.Errorf("want=%v get=%v", tt.want, res)
		}
	}
}

func TestSmallestDistancePair(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 3, 1}, 1, 0},
		{[]int{62, 100, 4}, 2, 58},
	}
	for _, tt := range tests {
		res := smallestDistancePair(tt.nums, tt.k)
		if tt.want != res {
			t.Errorf("want=%v get=%v", tt.want, res)
		}
	}
}

func TestBisectPeak(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 0},
		{[]int{2, 3, 1}, 2},
		{[]int{2, 3, 4, 5, 1}, 4},
	}
	for _, tt := range tests {
		res := bisectPeakL(tt.nums)
		if tt.want != res {
			t.Errorf("want=%v get=%v", tt.want, res)
		}
	}
}
