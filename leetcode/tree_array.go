package leetcode

import "sort"

type TreeArray struct {
	tree []int
}

func NewTreeArray(n int) *TreeArray {
	return &TreeArray{
		tree: make([]int, n+1),
	}
}

func lowbit(x int) int {
	return x & (-x)
}

func (t *TreeArray) update(k, d int) {
	for k <= len(t.tree) {
		t.tree[k] += d
		k += lowbit(k)
	}
}

func (t *TreeArray) sum(k int) (ans int) {
	for k > 0 {
		ans += t.tree[k]
		k -= lowbit(k)
	}
	return ans
}

func reversePairs(nums []int) (cnt int) {
	n := len(nums)
	hs0 := make(map[int]int, n*2)
	for _, c := range nums {
		hs0[c] = 0
		hs0[2*c] = 0
	}
	tmp := make([]int, len(hs0))
	for k := range hs0 {
		tmp = append(tmp, k)
	}
	sort.Ints(tmp)
	hs := make(map[int]int, len(tmp))
	k := 0
	for _, n := range tmp {
		k++
		hs[n] = k
	}
	ta := NewTreeArray(k)
	for _, a := range nums {
		cnt += ta.sum(k) - ta.sum(hs[2*a])
		ta.update(hs[a], 1)
	}
	return
}

func countRangeSum(nums []int, lower, upper int) (cnt int) {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	hs0 := make(map[int]int, n*3)
	for _, s := range preSum {
		hs0[s] = 0
		hs0[s-lower] = 0
		hs0[s-upper-1] = 0
	}
	tmp := make([]int, 0, len(hs0))
	for k := range hs0 {
		tmp = append(tmp, k)
	}
	sort.Ints(tmp)
	hs := make(map[int]int, len(tmp))
	k := 0
	for _, n := range tmp {
		k++
		hs[n] = k
	}

	ta := NewTreeArray(k)
	for i := 0; i <= n; i++ {
		s := preSum[i]
		r := hs[s-lower]
		l := hs[s-upper-1]
		cnt += ta.sum(r) - ta.sum(l)
		ta.update(hs[s], 1)
	}
	return
}
