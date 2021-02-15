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

func TestCountTriplets(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{2, 1, 3}, 12},
	}
	for _, tt := range tests {
		ret := countTripletsBit(tt.A)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestCanIWin(t *testing.T) {
	tests := []struct{
		chose int
		total int
		want bool
	}{
		{10, 11, false},
		{10, 10, true},
		{10, 40, false},
		{19, 190, true},
	}
	for _, tt := range tests {
		ret := canIWin(tt.chose, tt.total)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}