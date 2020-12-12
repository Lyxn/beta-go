package leetcode

import "testing"

func TestFindTheLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"lllll", 5},
		{"ele", 3},
		{"ellelel", 6},
		{"leetcodeisgreat", 5},
	}
	for _, tt := range tests {
		ret := findTheLongestSubstring(tt.s)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
