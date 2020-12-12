package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestQueen(t *testing.T) {
	res := solveNQueens(4)
	for i, r := range res {
		fmt.Printf("queen=%v\n%v\n", i, strings.Join(r, "\n"))
	}
}

func TestExistWord(t *testing.T) {
	tests := []struct {
		mat  [][]byte
		word string
		want bool
	}{
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCCE",
			true,
		},
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCB",
			false,
		},
	}
	for _, tt := range tests {
		ret := existWord(tt.mat, tt.word)
		if ret != tt.want {
			t.Errorf("word=%v want=%v get=%v", tt.word, tt.want, ret)
		}
	}
}

func TestMaxEnvelopes(t *testing.T) {
	tests := []struct {
		envStr string
		want   int
	}{
		{"[[46,89],[50,53],[52,68],[72,45],[77,81]]", 3},
	}
	for _, tt := range tests {
		env := make([][]int, 0)
		DecodeJson(tt.envStr, &env)
		ret := maxEnvelopes(env)
		if ret != tt.want {
			t.Logf("wamt=%v get=%v", tt.want, ret)
		}
	}
}
