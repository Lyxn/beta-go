package leetcode

import "math"

func getMidNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var last *ListNode
	for head.Next != nil {
		next := head.Next
		head.Next = last
		last = head
		head = next
	}
	head.Next = last
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	fake := &ListNode{Val: math.MinInt32, Next: nil}
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = nil
		fake = insert(fake, cur)
	}
	return fake.Next
}

func insert(lst, cur *ListNode) *ListNode {
	tmp := lst
	for tmp.Next != nil && tmp.Next.Val < cur.Val {
		tmp = tmp.Next
	}
	cur.Next = tmp.Next
	tmp.Next = cur
	return lst
}
