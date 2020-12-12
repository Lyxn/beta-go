package leetcode

import (
	"math"
	"sort"
)

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	tails := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		if v > tails[len(tails)-1] {
			tails = append(tails, v)
			continue
		}
		l := 0
		r := len(tails) - 1
		for l < r {
			m := (l + r) / 2
			if tails[m] >= v {
				r = m
			} else {
				l = m + 1
			}
		}
		tails[l] = v
	}
	return len(tails)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) (res float64) {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	} else if n2 == 0 {
		return 0
	} else if n1 == 0 {
		m := n2 / 2
		if n2%2 == 0 {
			return float64(nums2[m]+nums2[m-1]) / 2
		} else {
			return float64(nums2[m])
		}
	}
	n3 := n1 + n2
	l := 0
	r := n1
	m0 := 0
	m1 := 0
	for l <= r {
		x := (l + r) / 2
		y := (n3+1)/2 - x
		a0 := math.MinInt32
		if x != 0 {
			a0 = nums1[x-1]
		}
		a1 := math.MaxInt32
		if x != n1 {
			a1 = nums1[x]
		}
		b0 := math.MinInt32
		if y != 0 {
			b0 = nums2[y-1]
		}
		b1 := math.MaxInt32
		if y != n2 {
			b1 = nums2[y]
		}
		if b0 > a1 {
			l = x + 1
		} else if a0 > b1 {
			r = x - 1
		} else {
			m0 = max(a0, b0)
			m1 = min(a1, b1)
			break
		}
	}
	if n3%2 == 0 {
		return float64(m0+m1) / 2
	} else {
		return float64(m0)
	}
}

func bisect1(nums []int, k int) (idx int) {
	n := len(nums)
	l := 0
	r := n - 1
	if nums[l] > k {
		return 0
	} else if nums[r] <= k {
		return n
	}
	for l < r {
		m := (l + r) / 2
		if nums[m] > k {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	l := 0
	r := nums[n-1] - nums[0]
	for l < r {
		m := (l + r) / 2
		cnt := countPairs(nums, m)
		if cnt >= k {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func countPairs0(nums []int, w int) (cnt int) {
	n := len(nums)
	for r := 1; r < n; r++ {
		l := 0
		gap := nums[r] - w
		if nums[l] >= gap {
			cnt += r - l
			continue
		}
		h := r
		for l < h {
			m := (l + h) / 2
			if nums[m] < gap {
				l = m + 1
			} else {
				h = m
			}
		}
		cnt += r - l
	}
	return
}

func countPairs(nums []int, w int) (cnt int) {
	n := len(nums)
	l := 0
	for r := 1; r < n; r++ {
		for nums[r]-nums[l] > w {
			l++
		}
		cnt += r - l
	}
	return
}

func bisectPeakL(nums []int) int {
	n := len(nums)
	l := 0
	r := n - 1
	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func bisectPeakR(nums []int) int {
	n := len(nums)
	l := 0
	r := n - 1
	for l < r {
		m := (l + r + 1) / 2
		if nums[m] > nums[l] {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}
