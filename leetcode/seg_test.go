package leetcode

import "testing"

func TestGetSkyline(t *testing.T) {
	tests := []struct {
		buildings [][]int
	}{
		//{[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}},
		//{[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}}},
		//{[][]int{{3, 10, 20}, {3, 9, 19}, {3, 8, 18}, {3, 7, 17}, {3, 6, 16}, {3, 5, 15}, {3, 4, 14}}},
		{[][]int{{2, 4, 70}, {3, 8, 30}, {6, 100, 41}, {7, 15, 70}, {10, 30, 102}, {15, 25, 76}, {60, 80, 91}, {70, 90, 72}, {85, 120, 59}}},
	}
	for _, tt := range tests {
		res := getSkyline(tt.buildings)
		t.Logf("sky=%v", res)
	}
}

func TestInsertInterval(t *testing.T) {
	//intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	//newInterval := []int{4, 8}
	//intervals := [][]int{{1, 3}, {6, 9}}
	//newInterval := []int{2, 5}
	//intervals := [][]int{{1, 5}}
	//newInterval := []int{0, 0}
	intervals := [][]int{{3, 5}, {12, 15}}
	newInterval := []int{6, 6}

	ret := findInterval(intervals, newInterval[0])
	t.Logf("get=%v", ret)
	ret = findInterval(intervals, newInterval[1])
	t.Logf("get=%v", ret)
	res := insertInterval(intervals, newInterval)
	t.Logf("res=%v", res)
}
