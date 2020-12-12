package leetcode

import "testing"

func TestReverseList(t *testing.T) {
	head := MockListNode([]int{1, 2, 3})
	ret := reverseList(head)
	t.Logf("ret=%v", List2Ints(ret))
}

func TestInsertionSortList(t *testing.T) {
	head := MockListNode([]int{4, 1, 2, 3})
	ret := insertionSortList(head)
	t.Logf("ret=%v", List2Ints(ret))
}
