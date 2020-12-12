package leetcode

func TrapStack(height []int) int {
	stk := []int{}
	ret := 0
	for i := 0; i < len(height); i++ {
		for len(stk) > 0 && height[i] > height[stk[len(stk)-1]] {
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if len(stk) == 0 {
				break
			}
			top2 := stk[len(stk)-1]
			dist := i - top2 - 1
			hb := MinInt(height[i], height[top2]) - height[top]
			ret += dist * hb
		}
		stk = append(stk, i)
	}
	return ret
}

func TrapDP(height []int) (ret int) {
	n := len(height)
	if n <= 1 {
		return 0
	}
	lo := 0
	hi := n - 1
	lMax := height[lo]
	rMax := height[hi]
	for lo < hi {
		for lo < hi && lMax <= rMax {
			lo++
			if lMax < height[lo] {
				lMax = height[lo]
			} else {
				ret += lMax - height[lo]
			}
		}
		for lo < hi && rMax <= lMax {
			hi--
			if rMax < height[hi] {
				rMax = height[hi]
			} else {
				ret += rMax - height[hi]
			}
		}
	}
	return ret
}
