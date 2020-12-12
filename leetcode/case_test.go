package leetcode

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	res := generateParenthesis(4)
	t.Logf("get=%v\n", res)
}

func TestTrapDP(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		//{[]int{1, 0, 1}, 1},
		//{[]int{0, 1, 0, 2, 3, 1, 5}, 3},
		//{[]int{0, 1, 0, 2, 3, 1, 0, 5, 3}, 6},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}
	for i, tt := range tests {
		ret := TrapDP(tt.nums)
		if ret != tt.want {
			t.Errorf("idx=%v get=%v want=%v", i, ret, tt.want)
		}
	}
}
