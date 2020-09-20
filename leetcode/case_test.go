package leetcode

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	res := generateParenthesis(4)
	t.Logf("get=%v\n", res)
}

func mockListNode(nums []int) (res *ListNode) {
	tmp := res
	for _, val := range nums {
		node := &ListNode{Val: val}
		if tmp == nil {
			tmp = node
			res = tmp
		} else {
			tmp.Next = node
			tmp = tmp.Next
		}
	}
	return res
}

func list2Slice(node *ListNode) (res []int) {
	for {
		if node == nil {
			break
		}
		res = append(res, node.Val)
		node = node.Next
	}
	return
}

func TestAddTwoNum(t *testing.T) {
	a := mockListNode([]int{9})
	b := mockListNode([]int{9})
	c := addTwoNumbers(a, b)
	res := list2Slice(c)
	t.Logf("get=%v\n", res)
}

func TestThreeNum(t *testing.T) {
	//a := []int{1, -1, -1, 0}
	//a := []int{-1,0,1,2,-1,-4}
	a := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}

	res := threeSum(a)
	t.Logf("get=%v\n", res)
}

func TestValidParen(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{str: "{}[]", want: true},
		{str: "({}[])", want: true},
		{str: ")[]", want: false},
		{str: "[)]", want: false},
	}
	for _, tt := range tests {
		ret := isValidParen(tt.str)
		if ret != tt.want {
			t.Logf("str=%v want=%v get=%v\n", tt.str, tt.want, ret)
		}
	}
}

func TestGetTopK(t *testing.T) {
	tests := []struct{
		nums []int
		k int
		want int
	}{
		{
			nums: []int{1,2,1,3,4},
			k: 5,
			want: 4,
		},
	}
	for _, tt := range tests {
		ret := GetTopK(tt.nums, tt.k - 1)
		if tt.want != ret {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct{
		nums []int
		want []int
	}{
		{
			nums: []int{1,2,1,3,4},
			want: []int{1,1,2,3,4},
		},
	}
	for _, tt := range tests {
		QuickSort(tt.nums)
		t.Logf("get=%v want=%v", tt.nums, tt.want)
	}
}

func TestSmallestK(t *testing.T) {
	tests := []struct{
		nums []int
		k int
		want []int
	}{
		{
			nums: []int{1,2,1,3,4},
			k: 3,
			want: []int{1,1,2},
		},
	}
	for _, tt := range tests {
		ret := smallestK(tt.nums, tt.k)
		t.Logf("get=%v want=%v", ret, tt.want)
	}
}
