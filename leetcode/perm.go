package leetcode

import (
	"sort"
)

func permute(nums []int) [][]int {
	var res [][]int
	res = dfsPerm(res, nums, 0, 0)
	return res
}

func permuteUniq(nums []int) [][]int {
	var res [][]int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res = dfsPermUniq(res, nums, 0, 0)
	return res
}

func swap(path []int, i, j int) (res []int) {
	for _, p := range path {
		res = append(res, p)
	}
	res[i], res[j] = res[j], res[i]
	return
}

func dfsPerm(res [][]int, path []int, start int, depth int) (ret [][]int) {
	ret = res
	n := len(path)
	if depth == n {
		ret = append(ret, path)
		return ret
	}
	//fmt.Printf("start=%v depth=%v path=%v\n", start, depth, path)
	for i := start; i < n; i++ {
		tmp := swap(path, i, start)
		ret = dfsPerm(ret, tmp, start+1, depth+1)
	}
	return ret
}

func dfsPermUniq(res [][]int, path []int, start int, depth int) (ret [][]int) {
	ret = res
	n := len(path)
	if depth == n {
		ret = append(ret, path)
		return ret
	}
	//fmt.Printf("start=%v depth=%v path=%v\n", start, depth, path)
	vis := map[int]struct{}{}
	for i := start; i < n; i++ {
		if _, ok := vis[path[i]]; ok {
			continue
		}
		tmp := swap(path, i, start)
		ret = dfsPermUniq(ret, tmp, start+1, depth+1)
		vis[path[i]] = struct{}{}
	}
	return ret
}
