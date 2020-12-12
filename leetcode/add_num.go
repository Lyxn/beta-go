package leetcode

import (
	"sort"
)

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

func largestPerimeter(A []int) int {
	n := len(A)
	sort.Ints(A)
	res := 0
	l := 0
	r := 0
	for i := 0; i < n-2; i++ {
		l = max(l, i+1)
		r = max(r, l+1)
		if A[i]+A[l] <= A[r] {
			continue
		}
		for r < n {
			tl := l
			tr := r
			for tr < n && A[i]+A[tl] > A[tr] {
				tr++
			}
			tr--
			for tl < tr && A[i]+A[tl] > A[tr] {
				tl++
			}
			tl--
			if tr > tl {
				res = max(res, A[i]+A[tl]+A[tr])
			}
			if tr > r || tl > l {
				r = tr
				l = tl
			} else {
				break
			}
		}
	}
	return res
}

func largestPerimeter0(a []int) int {
	sort.Ints(a)
	for i := len(a) - 1; i >= 2; i-- {
		if a[i-2]+a[i-1] > a[i] {
			return a[i-2] + a[i-1] + a[i]
		}
	}
	return 0
}
