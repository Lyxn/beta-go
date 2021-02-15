package leetcode

import "testing"

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		net  [][]int
		want int
	}{
		{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
		{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
	}
	for _, tt := range tests {
		ret := findCircleNum(tt.net)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{
			[][]string{{"a", "b"}, {"b", "c"}},
			[]float64{2.0, 3.0},
			[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			[]float64{6, 0.5, -1, 1, -1},
		},
		{
			[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			[]float64{1.5, 2.5, 5.0},
			[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			[]float64{3.75000, 0.40000, 5.00000, 0.20000},
		},
	}
	for _, tt := range tests {
		ret := calcEquation(tt.equations, tt.values, tt.queries)
		if !isEqualFloats(ret, tt.want) {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
