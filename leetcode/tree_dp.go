package leetcode

func AddNode(root, cur *TreeNode) *TreeNode {
	if root == nil {
		return cur
	}
	if root.Val == cur.Val {
		root.Idx = cur.Idx
		return root
	}
	if root.Val > cur.Val {
		root.Left = AddNode(root.Left, cur)
	} else {
		root.Right = AddNode(root.Right, cur)
	}
	return root
}

func GetUB(root *TreeNode, dst int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == dst {
		return root
	}
	if dst > root.Val {
		return GetUB(root.Right, dst)
	}
	lub := GetUB(root.Left, dst)
	if lub == nil {
		return root
	}
	return lub
}

func GetLB(root *TreeNode, dst int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == dst {
		return root
	}
	if dst < root.Val {
		return GetLB(root.Left, dst)
	}
	rlb := GetLB(root.Right, dst)
	if rlb == nil {
		return root
	}
	return rlb
}

func oddEvenJumps(A []int) int {
	n := len(A)
	arc1 := make([]int, n)
	arc2 := make([]int, n)
	root := &TreeNode{Idx: n - 1, Val: A[n-1]}
	for i := n - 2; i >= 0; i-- {
		lb := GetLB(root, A[i])
		if lb != nil {
			arc2[i] = lb.Idx
		}
		ub := GetUB(root, A[i])
		if ub != nil {
			arc1[i] = ub.Idx
		}
		if ub != nil && ub.Val == A[i] {
			ub.Idx = i
		} else {
			root = AddNode(root, &TreeNode{Idx: i, Val: A[i]})
		}
	}

	dp1 := make([]int, n)
	dp2 := make([]int, n)
	dp1[n-1] = 1
	dp2[n-1] = 1
	cnt := 1
	for i := n - 2; i >= 0; i-- {
		n1 := arc1[i]
		n2 := arc2[i]
		if dp2[n1] > 0 {
			cnt += 1
			dp1[i] = 1
		}
		if dp1[n2] > 0 {
			dp2[i] = 1
		}
	}
	return cnt
}
