package leetcode

import "testing"

func TestOddEvenJump(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2}, 2},
		{[]int{10, 13, 12, 14, 15}, 2},
		{[]int{2, 3, 1, 1, 4}, 3},
		{[]int{5, 1, 3, 4, 2}, 3},
		{[]int{5, 4, 3, 2, 1}, 1},
	}
	for _, tt := range tests {
		ret := oddEvenJumps(tt.nums)
		if tt.want != ret {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
