package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

func fmtInts(nums []int, sep string) string {
	rs := make([]string, len(nums))
	for i, n := range nums {
		rs[i] = strconv.Itoa(n)
	}
	return strings.Join(rs, sep)
}

func TestKMP_Build(t *testing.T) {
	kmp := &KMP{}
	p := "aaabcaaabaaac"
	kmp.Build(p)
	t.Logf("\nkmp=%v\nnxt=%v\n", p, fmtInts(kmp.nxt, ""))
}

func TestKMP_Search(t *testing.T) {
	tests := []struct {
		p    string
		s    string
		want int
	}{
		{"ab", "abc", 0},
		{"aa", "abcaav", 3},
		{"abab", "abcabcabb", -1},
		{"abcab", "ababcabb", 2},
		{"aaabcaaabaaac", "dfadsfeaaaabcaaabaaac", 8},
	}
	kmp := &KMP{}
	for _, tt := range tests {
		kmp.Build(tt.p)
		ret := kmp.Search(tt.s)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
