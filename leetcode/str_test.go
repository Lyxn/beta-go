package leetcode

import (
	"testing"
)

func TestValidParen(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{str: "{}[]", want: true},
		{str: "({}[])", want: true},
		{str: ")[]", want: false},
		{str: "[)]", want: false},
	}
	for _, tt := range tests {
		ret := isValidParen(tt.str)
		if ret != tt.want {
			t.Logf("str=%v want=%v get=%v\n", tt.str, tt.want, ret)
		}
	}
}

func TestEditDist(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"horse", "ros", 3},
	}
	for _, tt := range tests {
		ret := EditDist(tt.a, tt.b)
		if ret != tt.want {
			t.Logf("str=%v %v want=%v get=%v\n", tt.a, tt.b, tt.want, ret)
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		a    string
		want int
	}{
		//{"", 0},
		//{"bbbb", 1},
		//{"horse", 5},
		//{"hooooo", 2},
		{"hooabooob", 3},
		//{"ooob", 2},
	}
	for _, tt := range tests {
		ret := LengthOfLongestSubstring(tt.a)
		if ret != tt.want {
			t.Errorf("str=%v want=%v get=%v\n", tt.a, tt.want, ret)
		}
	}
}

func TestLongestDupSubstring(t *testing.T) {
	tests := []struct {
		ss   string
		want string
	}{
		{"anana", "ana"},
		{"abcd", ""},
	}
	for _, tt := range tests {
		ret := longestDupSubstring(tt.ss)
		if ret != tt.want {
			t.Errorf("want=%v get=%v\n", tt.want, ret)
		}
	}
}

func TestShortestWay(t *testing.T) {
	tests := []struct {
		src  string
		dst  string
		want int
	}{
		{"abc", "abcbc", 2},
		{"xyz", "xzyxz", 3},
	}
	for _, tt := range tests {
		ret := shortestWay(tt.src, tt.dst)
		if ret != tt.want {
			t.Errorf("want=%v get=%v\n", tt.want, ret)
		}
	}
}

func TestMinCut(t *testing.T) {
	tests := []struct {
		ss   string
		want int
	}{
		{"fff", 0},
		{"ccaacabacb", 3},
		{"fifgbeajcacehiicccfecbfhhgfiiecdcjjffbghdidbhbdbfbfjccgbbdcjheccfbhafehieabbdfeigbiaggchaeghaijfbjhi", 75},
		{"adabdcaebdcebdcacaaaadbbcadabcbeabaadcbcaaddebdbddcbdacdbbaedbdaaecabdceddccbdeeddccdaabbabbdedaaabcdadbdabeacbeadbaddcbaacdbabcccbaceedbcccedbeecbccaecadccbdbdccbcbaacccbddcccbaedbacdbcaccdcaadcbaebebcceabbdcdeaabdbabadeaaaaedbdbcebcbddebccacacddebecabccbbdcbecbaeedcdacdcbdbebbacddddaabaedabbaaabaddcdaadcccdeebcabacdadbaacdccbeceddeebbbdbaaaaabaeecccaebdeabddacbedededebdebabdbcbdcbadbeeceecdcdbbdcbdbeeebcdcabdeeacabdeaedebbcaacdadaecbccbededceceabdcabdeabbcdecdedadcaebaababeedcaacdbdacbccdbcece", 273},
		{"bbdcdeaabdbaba", 5},
	}
	for _, tt := range tests {
		ret := minCut(tt.ss)
		if ret != tt.want {
			t.Errorf("want=%v get=%v\n", tt.want, ret)
		}
	}
}
