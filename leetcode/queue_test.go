package leetcode

import (
	"bytes"
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{
			grid: [][]byte{},
			want: 0,
		},
		{
			grid: [][]byte{
				{0, 1, 0},
				{1, 0, 1},
				{0, 1, 0},
			},
			want: 4,
		},
	}
	for i, tt := range tests {
		ret := numIslands(tt.grid)
		if tt.want != ret {
			t.Errorf("name=%v want=%v ret=%v", i, tt.want, ret)
		}
		t.Logf("ret=%v", ret)
	}
}

func TestSqrtInt(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 1},
		{12, 3},
	}
	for _, tt := range tests {
		ret := sqrtInt(tt.n)
		if ret != tt.want {
			t.Errorf("num=%v want=%v get=%v", tt.n, tt.want, ret)
		}
	}
}

func TestStr(t *testing.T) {
	s := "1"
	a := 'a'
	c := rune(s[0] + 1)
	buf := bytes.Buffer{}
	buf.WriteString(s[:0])
	buf.WriteRune(a)
	buf.WriteRune(c)
	buf.WriteString("end")
	t.Logf(buf.String())
}

func TestOpenLock(t *testing.T) {
	tests := []struct {
		deads  []string
		target string
		want   int
	}{
		{[]string{}, "8888", 8},
		{
			[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
			"8888",
			-1,
		},
	}
	for i, tt := range tests {
		ret := openLock(tt.deads, tt.target)
		if ret != tt.want {
			t.Errorf("num=%v want=%v get=%v", i, tt.want, ret)
		}
	}
}

func TestDecodeString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, tt := range tests {
		ret := decodeString(tt.s)
		if ret != tt.want {
			t.Errorf("in=%v want=%v get=%v", tt.s, tt.want, ret)
		}
	}
}

func TestFloodFill(t *testing.T) {
	tests := []struct {
		image [][]int
		sr    int
		sc    int
		color int
	}{
		{
			[][]int{{1, 1, 1}, {1, 1, 1}, {0, 1, 1}},
			1,
			1,
			2,
		},
	}
	for _, tt := range tests {
		ret := floodFill(tt.image, tt.sr, tt.sc, tt.color)
		t.Logf("get=%v", ret)
	}
}

func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
	}{
		//{
		//	[][]int{{0}, {1}},
		//},
		//{
		//	[][]int{{0, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1}},
		//},
		{
			[][]int{{1, 0, 1, 1, 0, 0, 1, 0, 0, 1}, {0, 1, 1, 0, 1, 0, 1, 0, 1, 1}, {0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 1, 0, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 1, 1, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0, 1, 0}, {1, 1, 1, 1, 0, 1, 0, 0, 1, 1}},
		},
	}
	for _, tt := range tests {
		ret := updateMatrix(tt.matrix)
		PrintMatrix(ret)
	}
}
