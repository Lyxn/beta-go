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

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want int
	}{
		{"ab", "b", 1},
		{"rabbbit", "rabbit", 3},
		{"babgbag", "bag", 5},
	}
	for _, tt := range tests {
		ret := numDistinct(tt.s1, tt.s2)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestIsEqualBase(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"eat", "aet", true},
		{"great", "rgaet", true},
		{"eatt", "aeta", false},
	}
	for _, tt := range tests {
		ret := isEqualBase(tt.s1, tt.s2)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestIsScramble(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"ab", "ba", true},
		{"great", "rgaet", true},
		{"abcde", "caebd", false},
		{"abcdbdac", "bdacabcd", true},
		{"abcdbdacbdac", "bdacabcdbdac", true},
	}
	for _, tt := range tests {
		ret := isScramble(tt.s1, tt.s2)
		if ret != tt.want {
			t.Errorf("get=%v want=%v s1=%v", ret, tt.want, tt.s1)
		}
	}
}

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		s   string
		want int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
	}
	for _, tt := range tests {
		ret := firstUniqChar(tt.s)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
