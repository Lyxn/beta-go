package leetcode

import "testing"

func TestMinimumDeleteSum(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want int
	}{
		{"sea", "eat", 231},
		{"delete", "leet", 403},
	}
	for _, tt := range tests {
		ret := minimumDeleteSum(tt.s1, tt.s2)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestSplitIntoFibonacci(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		//{"112", []int{1, 1, 2}},
		{"11235813", []int{1, 1, 2, 3, 5, 8, 13}},
	}
	for _, tt := range tests {
		ret := splitIntoFibonacci(tt.s)
		if !isEqualInts(ret, tt.want) {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{"a", "b", "ab", true},
		{"aabcc", "dbbca", "aadbbcbcac", true},
	}
	for _, tt := range tests {
		ret := isInterleave(tt.s1, tt.s2, tt.s3)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
