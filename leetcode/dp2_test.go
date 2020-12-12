package leetcode

import "testing"

func TestMinSwap(t *testing.T) {
	tests := []struct {
		as   []int
		bs   []int
		want int
	}{
		//{
		//	[]int{0, 4, 5},
		//	[]int{0, 1, 4},
		//	0,
		//},
		{
			[]int{0, 4, 4, 5, 9},
			[]int{0, 1, 6, 8, 10},
			1,
		},
	}
	for _, tt := range tests {
		ret := minSwap(tt.as, tt.bs)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestMinRefuelStops(t *testing.T) {
	tests := []struct {
		stations  [][]int
		target    int
		startFuel int
		want      int
	}{
		{nil, 1, 1, 0},
		{[][]int{{10, 100}}, 100, 1, -1},
		{
			[][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}},
			100, 10, 2,
		},
	}
	for _, tt := range tests {
		ret := minRefuelStops(tt.target, tt.startFuel, tt.stations)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestLargestSumOfAverages(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{9, 1, 2, 3, 9}, 1, 4.8},
		{[]int{9, 1, 2, 3, 9}, 3, 20},
	}
	for _, tt := range tests {
		ret := largestSumOfAverages(tt.nums, tt.k)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestProbabilityOfHeads(t *testing.T) {
	tests := []struct {
		prob   []float64
		target int
		want   float64
	}{
		{[]float64{0.2, 0.8, 0.3}, 3, 0.048},
		{[]float64{0.2, 0.8, 0.3, 0.5}, 3, 0.182},
	}
	for _, tt := range tests {
		ret := probabilityOfHeads(tt.prob, tt.target)
		if !isEqualFloat(ret, tt.want) {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestMinDistance(t *testing.T) {
	tests := []struct {
		houses []int
		k      int
		want   int
	}{
		{[]int{2, 5, 7}, 1, 5},
		{[]int{1, 4, 10}, 1, 9},
		{[]int{2, 3, 5, 12, 18}, 2, 9},
		{[]int{2, 5, 7, 10, 14}, 2, 9},
		{[]int{1, 4, 10, 14, 18, 19, 20, 26, 28}, 3, 18},
	}
	for _, tt := range tests {
		ret := minDistance(tt.houses, tt.k)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}
