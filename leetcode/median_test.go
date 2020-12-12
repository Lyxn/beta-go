package leetcode

import "testing"

func TestNewMedianFinder(t *testing.T) {
	mf := NewMedianFinder()
	nums := []int{-1, -2, -3, -4, -5}
	wants := []float64{-1, -1.5, -2, -2.5, -3}
	for i, n := range nums {
		mf.AddNum1(n)
		ret := mf.FindMedian()
		if wants[i] != ret {
			t.Errorf("add=%v get=%v want=%v", n, ret, wants[i])
		}
	}
}
