package leetcode

import (
	"fmt"
	"math"
	"sync"
)

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

func HeapSort(nums []int) {
	heapify(nums)
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		adjustHeap(nums, 0, i)
	}
}

func heapify(nums []int) {
	n := len(nums)
	for p := n - 2; p >= 0; p-- {
		adjustHeap(nums, p, n)
	}
}

func adjustHeap(nums []int, p, n int) {
	l := p*2 + 1
	r := p*2 + 2
	maxIdx := p
	maxVal := nums[p]
	for l < n {
		if nums[l] > maxVal {
			maxVal = nums[l]
			maxIdx = l
		}
		if r < n && nums[r] > maxVal {
			maxVal = nums[r]
			maxIdx = r
		}
		if maxIdx == p {
			break
		}
		nums[maxIdx], nums[p] = nums[p], nums[maxIdx]
		p = maxIdx
		maxVal = nums[maxIdx]
		l = p*2 + 1
		r = p*2 + 2
	}
}

func topKFrequent(nums []int, k int) []int {
	dct := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dct[nums[i]] += 1
	}
	cnt := make([][2]int, 0, len(dct))
	for k, c := range dct {
		cnt = append(cnt, [2]int{k, c})
	}
	return getTopKCnt(cnt, k)
}

func adjustCnt(cnt [][2]int, s, e int) {
	i := s
	mi := i
	mv := cnt[i][1]
	l := 2*i + 1
	r := 2*i + 2
	for l < e {
		if cnt[l][1] > mv {
			mi = l
			mv = cnt[l][1]
		}
		if r < e && cnt[r][1] > mv {
			mi = r
			mv = cnt[r][1]
		}
		if mi == i {
			break
		}
		cnt[i], cnt[mi] = cnt[mi], cnt[i]
		i = mi
		mv = cnt[i][1]
		l = 2*i + 1
		r = 2*i + 2
	}
}

func getTopKCnt(cnt [][2]int, k int) []int {
	n := len(cnt)
	for j := n - 2; j >= 0; j-- {
		adjustCnt(cnt, j, n)
	}
	res := make([]int, k)
	res[0] = cnt[0][0]
	for i := 1; i < k; i++ {
		cnt[0], cnt[n-i] = cnt[n-i], cnt[0]
		adjustCnt(cnt, 0, n-i)
		res[i] = cnt[0][0]
	}
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	ln := make([]*ListNode, 0, len(lists))
	for _, nd := range lists {
		if nd != nil {
			ln = append(ln, nd)
		}
	}
	lists = ln
	if len(lists) == 0 {
		return nil
	}
	heapifyList(lists)
	root := &ListNode{}
	pre := root
	for len(lists) > 0 {
		adjustList(lists, 0, len(lists))
		cur := lists[0]
		pre.Next = cur
		pre = cur
		if cur.Next == nil {
			lists[0] = lists[len(lists)-1]
			lists = lists[:len(lists)-1]
		} else {
			lists[0] = cur.Next
		}
	}
	return root.Next
}

func adjustList(ln []*ListNode, s, e int) {
	mi := s
	mv := ln[s].Val
	l := 2*s + 1
	r := 2*s + 2
	for l < e {
		if ln[l].Val < mv {
			mv = ln[l].Val
			mi = l
		}
		if r < e && ln[r].Val < mv {
			mv = ln[r].Val
			mi = r
		}
		if mi == s {
			break
		}
		ln[s], ln[mi] = ln[mi], ln[s]
		s = mi
		mv = ln[s].Val
		l = 2*s + 1
		r = 2*s + 2
	}
}

func heapifyList(ln []*ListNode) {
	n := len(ln)
	for i := n - 2; i >= 0; i-- {
		adjustList(ln, i, n)
	}
}

func mergeKListsGo(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var wg sync.WaitGroup
	for len(lists) >= 2 {
		newLst := []*ListNode{}
		n := len(lists)
		if n%2 == 1 {
			newLst = append(newLst, lists[0])
			lists = lists[1:]
		}
		chlst := make(chan *ListNode, 1)
		for i := 0; i < n/2; i++ {
			wg.Add(1)
			go func(x int) {
				defer wg.Done()
				rl := mergeList(lists[2*x], lists[2*x+1])
				chlst <- rl
			}(i)
		}
		go func() {
			wg.Wait()
			close(chlst)
		}()
		for lst := range chlst {
			newLst = append(newLst, lst)
		}
		lists = newLst
	}
	return lists[0]
}

func mergeList(a, b *ListNode) *ListNode {
	root := &ListNode{}
	pre := root
	for a != nil && b != nil {
		if a.Val < b.Val {
			pre.Next = a
			a = a.Next
		} else {
			pre.Next = b
			b = b.Next
		}
		pre = pre.Next
	}
	if a != nil {
		pre.Next = a
	} else {
		pre.Next = b
	}
	return root.Next
}

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	maxVal := math.MinInt32
	minVal := math.MaxInt32
	for i := 0; i < n; i++ {
		maxVal = max(maxVal, nums[i])
		minVal = min(minVal, nums[i])
	}
	maxGap := maxVal - minVal
	if maxGap == 0 {
		return 0
	}
	avgGap := max(maxGap/(n-1), 1)
	numBck := maxGap/avgGap + 1
	maxBck := make([]int, numBck)
	minBck := make([]int, numBck)
	for i := 0; i < numBck; i++ {
		maxBck[i] = math.MinInt32
		minBck[i] = math.MaxInt32
	}
	for i := 0; i < n; i++ {
		bi := (nums[i] - minVal) / avgGap
		maxBck[bi] = max(maxBck[bi], nums[i])
		minBck[bi] = min(minBck[bi], nums[i])
	}
	pre := maxBck[0]
	res := avgGap
	for k := 1; k < numBck; k++ {
		if minBck[k] == math.MaxInt32 {
			continue
		}
		res = max(res, minBck[k]-pre)
		pre = maxBck[k]
	}
	return res
}
