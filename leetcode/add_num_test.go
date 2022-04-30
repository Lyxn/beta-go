package leetcode

import (
	"testing"

	"leetcode-go/util"
)

func TestAddTwoNum(t *testing.T) {
	a := MockListNode([]int{9})
	b := MockListNode([]int{9})
	c := addTwoNumbers(a, b)
	res := List2Ints(c)
	t.Logf("get=%v\n", res)
}

func TestThreeNum(t *testing.T) {
	//a := []int{1, -1, -1, 0}
	//a := []int{-1,0,1,2,-1,-4}
	a := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}
	res := threeSum(a)
	t.Logf("get=%v\n", res)
}

func TestLargestPerimeter(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 4, 4, 9, 15}, 11},
		{[]int{1, 2, 2, 4, 18, 8}, 5},
	}
	for _, tt := range tests {
		ret := largestPerimeter(tt.nums)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestLargestPerimeter1(t *testing.T) {
	filename := "./case/largest-perimeter-80.txt"
	nums := util.ReadInts(filename)
	ret := largestPerimeter(nums[0])
	t.Logf("len=%v get=%v", len(nums[0]), ret)
}
