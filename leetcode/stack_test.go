package leetcode

import "testing"

func TestRemoveKdigits(t *testing.T) {
	tests := []struct {
		num  string
		k    int
		want string
	}{
		{"1432219", 3, "1219"},
	}
	for _, tt := range tests {
		ret := removeKdigits(tt.num, tt.k)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestRemoveDuplicateLetters(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"abcab", "abc"},
		{"bcabc", "abc"},
		{"cbacdcbc", "acdb"},
		{"abacb", "abc"},
	}
	for _, tt := range tests {
		ret := removeDuplicateLetters(tt.s)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
