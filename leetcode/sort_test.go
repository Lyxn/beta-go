package leetcode

import (
	"sort"
	"testing"

	"leetcode-go/util"
)

func TestHeapSort(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{
			nums: []int{1, 2, 1, 3, 4},
		},
		{
			nums: []int{1, 2, 1, 5, 9, 10, 3, 4},
		},
	}
	for _, tt := range tests {
		HeapSort(tt.nums)
		t.Logf("get=%v", tt.nums)
	}
}

func TestSmallestK(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 2, 1, 3, 4},
			k:    3,
			want: []int{1, 1, 2},
		},
	}
	for _, tt := range tests {
		ret := smallestK(tt.nums, tt.k)
		t.Logf("get=%v want=%v", ret, tt.want)
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 1, 3, 4},
			want: []int{1, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		QuickSort(tt.nums)
		t.Logf("get=%v want=%v", tt.nums, tt.want)
	}
}

func TestGetTopK(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{1, 2, 1, 3, 4},
			k:    5,
			want: 4,
		},
	}
	for _, tt := range tests {
		ret := GetTopK(tt.nums, tt.k-1)
		if tt.want != ret {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		//{[]int{-1, 1, 4, -4, 3, 5, 4, -2, 3, -1}, 3, []int{-1, 3, 4}},
		{
			[]int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3},
			3,
			[]int{3, 7, 10},
		},
	}
	for i, tt := range tests {
		ret := topKFrequent(tt.nums, tt.k)
		sort.Ints(ret)
		if !isEqualInts(ret, tt.want) {
			t.Logf("idx=%v get=%v want=%v", i, ret, tt.want)
		}
	}
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		strLst string
		want   string
	}{
		{
			strLst: "[[-10,-9,-9,-3,-1,-1,0],[-5],[4],[-8],[],[-9,-6,-5,-4,-2,2,3],[-3,-3,-2,-1,0]]",
			want:   "[-10,-9,-9,-9,-8,-6,-5,-5,-4,-3,-3,-3,-2,-2,-1,-1,-1,0,0,2,3,4]",
		},
		{
			strLst: "[[1,2],[2,3]]",
			want:   "[1,2,2,3]",
		},
	}
	for _, tt := range tests {
		intList := make([][]int, 0)
		DecodeJson(tt.strLst, &intList)
		lst := make([]*ListNode, len(intList))
		for i, ints := range intList {
			lst[i] = MockListNode(ints)
		}
		ret := mergeKListsGo(lst)
		ints := List2Ints(ret)
		get := EncodeJson(ints)
		if get != tt.want {
			t.Errorf("want=%v get=%v", tt.want, get)
		}
	}
}

func TestMaximumGap(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 1, 1}, 0},
		{[]int{10, 1, 2, 9}, 7},
	}
	for _, tt := range tests {
		ret := maximumGap(tt.nums)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
	}
}

func TestMaximumGap1(t *testing.T) {
	filename := "./case/maximum-gap-15.txt"
	tests := util.ReadInts(filename)
	for _, tt := range tests {
		ret := maximumGap(tt)
		t.Logf("get=%v", ret)
	}
}
