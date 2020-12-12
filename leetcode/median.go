package leetcode

type MedianFinder struct {
	maxHp []int
	minHp []int
}

/** initialize your data structure here. */
func NewMedianFinder() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum1(num int) {
	sizeMaxHp := mf.getSizeOfMaxHp()
	sizeMinHp := mf.getSizeOfMinHp()
	if sizeMinHp == 0 && sizeMaxHp == 0 {
		mf.insertMaxHp(num)
		return
	}
	if sizeMinHp == 0 {
		if mf.maxHp[0] > num {
			mf.minHp = append(mf.minHp, mf.maxHp[0])
			mf.maxHp[0] = num
		} else {
			mf.minHp = append(mf.minHp, num)
		}
		return
	}
	if sizeMinHp == sizeMaxHp {
		if num > mf.minHp[0] {
			mf.insertMinHp(num)
		} else {
			mf.insertMaxHp(num)
		}
	} else if sizeMaxHp > sizeMinHp {
		if num < mf.maxHp[0] {
			mf.insertMinHp(mf.maxHp[0])
			mf.maxHp[0] = num
			dwMaxUp(mf.maxHp, 0, len(mf.maxHp))
		} else {
			mf.insertMinHp(num)
		}
	} else {
		if num > mf.minHp[0] {
			mf.insertMaxHp(mf.minHp[0])
			mf.minHp[0] = num
			dwMinUp(mf.minHp, 0, len(mf.minHp))
		} else {
			mf.insertMaxHp(num)
		}
	}
}

func (mf *MedianFinder) AddNum(num int) {
	sizeMaxHp := mf.getSizeOfMaxHp()
	sizeMinHp := mf.getSizeOfMinHp()
	if sizeMaxHp <= sizeMinHp {
		mf.insertMaxHp(num)
	} else {
		mf.insertMinHp(num)
	}
	if sizeMinHp == 0 && sizeMaxHp == 0 {
		return
	}
	for mf.maxHp[0] > mf.minHp[0] {
		mf.maxHp[0], mf.minHp[0] = mf.minHp[0], mf.maxHp[0]
		dwMaxUp(mf.maxHp, 0, len(mf.maxHp))
		dwMinUp(mf.minHp, 0, len(mf.minHp))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	sizeMaxHp := mf.getSizeOfMaxHp()
	sizeMinHp := mf.getSizeOfMinHp()
	if sizeMaxHp == 0 && sizeMinHp == 0 {
		return -1
	} else if sizeMinHp == 0 {
		return mf.getTopOfMaxHp()
	} else if sizeMaxHp == 0 {
		return mf.getTopOfMinHp()
	}
	if sizeMaxHp == sizeMinHp {
		return (mf.getTopOfMinHp() + mf.getTopOfMaxHp()) / 2.0
	} else if sizeMinHp > sizeMaxHp {
		return mf.getTopOfMinHp()
	} else {
		return mf.getTopOfMaxHp()
	}
}

func (mf *MedianFinder) getTopOfMaxHp() float64 {
	return float64(mf.maxHp[0])
}

func (mf *MedianFinder) getTopOfMinHp() float64 {
	return float64(mf.minHp[0])
}

func (mf *MedianFinder) getSizeOfMaxHp() int {
	return len(mf.maxHp)
}

func (mf *MedianFinder) getSizeOfMinHp() int {
	return len(mf.minHp)
}

func (mf *MedianFinder) insertMaxHp(num int) {
	mf.maxHp = append(mf.maxHp, num)
	upMaxHp(mf.maxHp, len(mf.maxHp)-1)
}

func (mf *MedianFinder) insertMinHp(num int) {
	mf.minHp = append(mf.minHp, num)
	upMinHp(mf.minHp, len(mf.minHp)-1)
}

func upMaxHp(nums []int, x int) {
	f := getParent(x)
	for f >= 0 {
		if nums[f] >= nums[x] {
			break
		}
		nums[x], nums[f] = nums[f], nums[x]
		x = f
		f = getParent(x)
	}
}

func upMinHp(nums []int, x int) {
	f := getParent(x)
	for f >= 0 {
		if nums[f] <= nums[x] {
			break
		}
		nums[x], nums[f] = nums[f], nums[x]
		x = f
		f = getParent(x)
	}
}

func dwMaxUp(nums []int, s, e int) {
	if len(nums) <= s {
		return
	}
	i := s
	l := 2*i + 1
	r := 2*i + 2
	for l < e {
		mi := i
		if nums[l] > nums[mi] {
			mi = l
		}
		if r < e && nums[r] > nums[mi] {
			mi = r
		}
		if mi == i {
			break
		}
		nums[i], nums[mi] = nums[mi], nums[i]
		i = mi
		l = 2*i + 1
		r = 2*i + 2
	}
}

func dwMinUp(nums []int, s, e int) {
	if len(nums) <= s {
		return
	}
	i := s
	l := 2*i + 1
	r := 2*i + 2
	for l < e {
		mi := i
		if nums[l] < nums[mi] {
			mi = l
		}
		if r < e && nums[r] < nums[mi] {
			mi = r
		}
		if mi == i {
			break
		}
		nums[i], nums[mi] = nums[mi], nums[i]
		i = mi
		l = 2*i + 1
		r = 2*i + 2
	}
}

func getParent(x int) int {
	if x == 0 {
		return -1
	}
	if x%2 == 0 {
		return (x - 2) / 2
	} else {
		return (x - 1) / 2
	}
}
