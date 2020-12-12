package leetcode

func wiggleMaxLength0(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	i := 0
	for i < n-1 && nums[i+1] == nums[i] {
		i++
	}
	up := nums[i+1] > nums[i]
	i += 1
	cnt := 2
	for ; i < n-1; i++ {
		if nums[i+1] == nums[i] {
			continue
		}
		if up {
			if nums[i+1] < nums[i] {
				up = !up
				cnt += 1
			}
		} else {
			if nums[i+1] > nums[i] {
				up = !up
				cnt += 1
			}
		}
	}
	return cnt
}

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	up := 1
	dw := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = dw + 1
		} else if nums[i] < nums[i-1] {
			dw = up + 1
		}
	}
	return max(up, dw)
}
