package leetcode

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	addTwo(res, l1, l2, 0)
	return res.Next
}

func addTwo(res, l1, l2 *ListNode, up int) {
	if l1 == nil && l2 == nil {
		if up != 0 {
			res.Next = &ListNode{Val: up, Next: nil}
		}
		return
	}
	val := 0
	nextUp := 0
	if l1 == nil {
		val = l2.Val + up
		l2 = l2.Next
	} else if l2 == nil {
		val = l1.Val + up
		l1 = l1.Next
	} else {
		val = l1.Val + l2.Val + up
		l1 = l1.Next
		l2 = l2.Next
	}
	if val > 9 {
		nextUp = val / 10
		val = val % 10
	}
	res.Next = &ListNode{Val: val, Next: nil}
	addTwo(res.Next, l1, l2, nextUp)
}

func threeSum(nums []int) (res [][]int) {
	length := len(nums)
	if length < 3 {
		return
	}
	sort.Ints(nums)
	lo := 0
	for lo < length {
		res2 := twoSum(nums, lo+1, length-1, -nums[lo])
		res = append(res, res2...)
		lo = getNextIdx(nums, lo)
	}
	return res
}

func getNextIdx(nums []int, idx int) (res int) {
	for idx = idx + 1; idx < len(nums); idx++ {
		if nums[idx-1] != nums[idx] {
			break
		}
	}
	return idx
}

func getLastIdx(nums []int, idx int) (res int) {
	for idx = idx - 1; idx >= 0; idx-- {
		if nums[idx+1] != nums[idx] {
			break
		}
	}
	return idx
}

func twoSum(nums []int, lo, hi int, target int) (res [][]int) {
	for lo < hi {
		if nums[lo]+nums[hi] == target {
			res = append(res, []int{-target, nums[lo], nums[hi]})
		}
		if nums[lo]+nums[hi] < target {
			lo = getNextIdx(nums, lo)
		} else {
			hi = getLastIdx(nums, hi)
		}
	}
	return
}
