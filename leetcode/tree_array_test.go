package leetcode

import (
	"testing"

	"leetcode-go/util"
)

func TestReversePair(t *testing.T) {
	nums := []int{7, 3, 8, 3, 1}
	res := reversePairs(nums)
	t.Logf("get=%v", res)
}

func TestReversePair1(t *testing.T) {
	filename := "./case/reverse-pairs-37.txt"
	nums := util.ReadInts(filename)
	res := reversePairs(nums[0])
	t.Logf("get=%v", res)
}

func TestCountRangeSum(t *testing.T) {
	tests := []struct {
		nums  []int
		lower int
		upper int
		want  int
	}{
		{[]int{-2, 5, -1}, -2, 2, 3},
	}
	for _, tt := range tests {
		res := countRangeSum(tt.nums, tt.lower, tt.upper)
		if res != tt.want {
			t.Errorf("get=%v want=%v", res, tt.want)
		}
	}
}
