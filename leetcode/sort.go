package leetcode

import "fmt"

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	base := nums[0]
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		for lo < hi && nums[hi] >= base {
			hi--
		}
		nums[lo] = nums[hi]
		for lo < hi && nums[lo] <= base {
			lo++
		}
		nums[hi] = nums[lo]
	}
	nums[lo] = base
	fmt.Printf("nums=%v\n", nums)
	QuickSort(nums[:lo])
	QuickSort(nums[lo+1:])
}

func GetTopK(nums []int, k int) (res int) {
	base := nums[0]
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		for lo < hi && nums[hi] >= base {
			hi--
		}
		nums[lo] = nums[hi]
		for lo < hi && nums[lo] <= base {
			lo++
		}
		nums[hi] = nums[lo]
	}
	nums[lo] = base
	fmt.Printf("%v %v\n", nums, k)
	if lo == k {
		return nums[lo]
	} else if lo > k {
		return GetTopK(nums[:lo], k)
	} else {
		return GetTopK(nums[lo+1:], k-lo-1)
	}
}

func smallestK(arr []int, k int) []int {
	num := len(arr)
	if k >= num {
		return arr
	} else if k == 0 {
		return nil
	}
	lo := 0
	hi := num - 1
	pos := part(arr, lo, hi)
	k -= 1
	for pos != k {
		if pos > k {
			hi = pos - 1
		} else {
			lo = pos + 1
		}
		pos = part(arr, lo, hi)
	}
	return arr[:pos+1]
}

func part(arr []int, beg, end int) int {
	base := arr[beg]
	for beg < end {
		for beg < end && arr[end] >= base {
			end--
		}
		arr[beg] = arr[end]
		for beg < end && arr[beg] <= base {
			beg++
		}
		arr[end] = arr[beg]
	}
	arr[beg] = base
	return beg
}
