package leetcode

import "testing"

func TestMinCostConnectPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{
			[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			20,
		},
		{
			[][]int{{2, -3}, {-17, -8}, {13, 8}, {-17, -15}},
			53,
		},
	}
	for _, tt := range tests {
		ret := minCostConnectPointsUF(tt.points)
		if ret != tt.want {
			t.Errorf("want=%v get=%v", tt.want, ret)
		}
	}
}
