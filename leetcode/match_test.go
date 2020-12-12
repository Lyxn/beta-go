package leetcode

import "testing"

func TestIsMatchWild(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want bool
	}{
		{"ab", "a*", true},
		{"mississippi", "m??*ss*?i*pi", false},
		{"b", "?*?", false},
		{
			"bbbaaabaababbabbbaabababbbabababaabbaababbbabbbabb",
			"*b**b***baba***aaa*b***",
			false,
		},
		{
			"abcabczzzde",
			"*abc???de*",
			true,
		},
		{"aaaa", "***a", true},
	}
	for _, tt := range tests {
		ret := isMatchWildGreedy(tt.s, tt.p)
		if ret != tt.want {
			t.Errorf("p=%v want=%v get=%v", tt.p, tt.want, ret)
		}
	}
}
