package leetcode

import "testing"

func TestGetPathDist(t *testing.T) {
	tests := []struct {
		A    []string
		st   []int
		want int
	}{
		{
			[]string{"catg", "ctaagt", "gcta", "ttca"},
			[]int{0, 1, 2, 3},
			18,
		},
		{
			[]string{"catg", "ctaagt", "gcta", "ttca"},
			[]int{2, 1, 3, 0},
			12,
		},
	}
	for _, tt := range tests {
		pr := buildOverlap(tt.A)
		res := getPathDist(tt.A, pr, tt.st)
		if res != tt.want {
			t.Errorf("want=%v get=%v", tt.want, res)
		}
	}
}

func TestShortestSuperstring(t *testing.T) {
	tests := []struct {
		A    []string
		want string
	}{
		{
			[]string{"loves", "leetcode"},
			"lovesleetcode",
		},
		{
			[]string{"catg", "ctaagt", "gcta", "ttca"},
			"gctaagttcatg",
		},
		{
			[]string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"},
			"gctaagttcatgcatc",
		},
	}
	for _, tt := range tests {
		res := shortestSuperstring(tt.A)
		//if len(res) != len(tt.want) {
		if res != tt.want {
			t.Errorf("want=%v get=%v", tt.want, res)
		}
	}
}
