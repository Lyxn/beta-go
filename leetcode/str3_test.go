package leetcode

import "testing"

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

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		ss   string
		want int
	}{
		{"aa", 3},
		//{"abc", 3},
	}
	for _, tt := range tests {
		ret := countSubstrings(tt.ss)
		if ret != tt.want {
			t.Errorf("want=%v get=%v\n", tt.want, ret)
		}
	}
}

func TestLongestDecomposition(t *testing.T) {
	tests := []struct {
		ss   string
		want int
	}{
		{"aa", 2},
		{"aaa", 3},
	}
	for _, tt := range tests {
		ret := longestDecomposition(tt.ss)
		if ret != tt.want {
			t.Errorf("want=%v get=%v\n", tt.want, ret)
		}
	}
}

func TestLongestPalindromeSubseq(t *testing.T) {
	tests := []struct {
		ss   string
		want int
	}{
		{"aa", 2},
		{"abda", 3},
	}
	for _, tt := range tests {
		ret := longestPalindromeSubseq(tt.ss)
		if ret != tt.want {
			t.Errorf("want=%v get=%v\n", tt.want, ret)
		}
	}
}

func TestStrangePrinter(t *testing.T) {
	tests := []struct {
		ss   string
		want int
	}{
		{"aaabbb", 2},
		{"aba", 2},
		{"abda", 3},
		{"abada", 3},
		{"bccbbaddbacaddabaddcdadbadbccaccdcbbdccbcd", 19},
		{"bcbadbacadabadcdadbadbcacdcbdcbcd", 19},
		{"dbadbcacdcbdcbcd", 9},
	}
	for _, tt := range tests {
		ret := strangePrinter(tt.ss)
		if ret != tt.want {
			t.Errorf("want=%v get=%v\n", tt.want, ret)
		}
	}
}

func TestRepeatedSubString(t *testing.T) {
	tests := []struct {
		ss   string
		want int
	}{
		{"aaabbb", -1},
		{"ababab", 1},
		{"abcabc", 2},
	}
	for _, tt := range tests {
		ret := repeatedSubString(tt.ss)
		if ret != tt.want {
			t.Errorf("want=%v get=%v\n", tt.want, ret)
		}
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		ss   string
		want string
	}{
		//{"aaabbb", "aaabbb"},
		//{"ababab", "3[ab]"},
		//{"abababc", "3[ab]c"},
		{"aaaaaaaaaa", "10[a]"},
	}
	for _, tt := range tests {
		ret := encode(tt.ss)
		if ret != tt.want {
			t.Errorf("want=%v get=%v\n", tt.want, ret)
		}
	}
}
