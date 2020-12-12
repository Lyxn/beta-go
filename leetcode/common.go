package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Idx   int
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MockListNode(ln []int) *ListNode {
	if len(ln) == 0 {
		return nil
	}
	root := &ListNode{Val: ln[0]}
	pre := root
	for i := 1; i < len(ln); i++ {
		pre.Next = &ListNode{Val: ln[i]}
		pre = pre.Next
	}
	return root
}

func List2Ints(node *ListNode) (res []int) {
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return
}
