package leetcode

import "testing"

func TestWiggleMaxLength(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 3, 4}, 2},
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
	}
	for _, tt := range tests {
		ret := wiggleMaxLength(tt.nums)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
