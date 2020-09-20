package leetcode

import (
	"fmt"
	"testing"
)

func TestPerm(t *testing.T) {
	a := []int{1, 2, 3}
	res := permute(a)
	fmt.Printf("%v\n", res)
}

func TestPermUniq(t *testing.T) {
	a := []int{0, 0, 1, 9}
	res := permuteUniq(a)
	for i := range res {
		fmt.Printf("%v\n", res[i])
	}
}
