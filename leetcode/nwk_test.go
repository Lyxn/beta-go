package leetcode

import "testing"

func TestLadderLength(t *testing.T) {
	beg := "hit"
	end := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	ret := ladderLength(beg, end, wordList)
	t.Logf("get=%v", ret)
}

func TestAlienOrder(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		//{[]string{"abc"}, "abc"},
		//{[]string{"abc", "ab"}, ""},
		//{[]string{"ac", "ab"}, "cab"},
		//{[]string{"z", "x"}, "zx"},
		//{[]string{"z", "x", "z"}, ""},
		{[]string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
	}
	for _, tt := range tests {
		ret := alienOrder(tt.words)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
