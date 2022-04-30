package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	NilStr  string = "null"
	NilNode int    = math.MaxInt32
)

func (t *TreeNode) AddNode(nd *TreeNode) {
	if t.Val > nd.Val {
		if t.Right == nil {
			t.Right = nd
		}
		t.Right.AddNode(nd)
	} else if t.Val < nd.Val {
		if t.Left == nil {
			t.Left = nd
		}
		t.Left.AddNode(nd)
	}
}

func BuildTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	for i := 1; i < n; i++ {
		nd := &TreeNode{Val: nums[i]}
		root.AddNode(nd)
	}
	return root
}

func PrintNodes(nodes []*TreeNode) {
	for _, nd := range nodes {
		fmt.Printf("node=%p %+v\n", nd, nd)
	}
}

func BuildBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	n := len(nums)
	nodes := make([]*TreeNode, n)
	for i := 0; i < n; i++ {
		val := nums[i]
		if val != NilNode {
			nodes[i] = &TreeNode{Val: val}
		}
	}
	//PrintNodes(nodes)
	for i := 1; i < n; i++ {
		var p int
		if i%2 == 1 {
			p = (i - 1) / 2
			nodes[p].Left = nodes[i]
		} else {
			p = i/2 - 1
			nodes[p].Right = nodes[i]
		}
	}
	//PrintNodes(nodes)
	return nodes[0]
}

func serialize(root *TreeNode) (str string) {
	var strs []string
	nodes := []*TreeNode{root}
	for {
		n := len(nodes)
		hasNext := false
		for i := 0; i < n; i++ {
			nd := nodes[i]
			if nd == nil {
				strs = append(strs, NilStr)
			} else {
				strs = append(strs, strconv.Itoa(nd.Val))
				nodes = append(nodes, nd.Left, nd.Right)
				hasNext = hasNext || nd.Left != nil || nd.Right != nil
			}
		}
		if !hasNext {
			break
		}
		nodes = nodes[n:]
	}
	i := len(strs) - 1
	for ; i >= 0; i-- {
		if strs[i] != NilStr {
			break
		}
	}
	strs = strs[:i+1]
	str = fmt.Sprintf("[%v]", strings.Join(strs, ","))
	return
}

func deserialize(str string) (root *TreeNode) {
	nb := len(str)
	if nb == 0 || str == "[]" {
		return nil
	}
	vals := strings.Split(str[1:nb-1], ",")
	n := len(vals)
	nodes := make([]*TreeNode, n)
	for i := 0; i < n; i++ {
		if vals[i] == NilStr {
			nodes[i] = nil
		} else {
			val, _ := strconv.Atoi(vals[i])
			nodes[i] = &TreeNode{Val: val}
		}
	}
	root = nodes[0]
	lvs := []*TreeNode{root}
	start := 0
	for len(lvs) > 0 {
		idx := start
		tmpN := len(lvs)
		for i := 0; i < tmpN; i++ {
			nd := lvs[i]
			idx += 1
			if idx < n {
				nd.Left = nodes[idx]
			}
			idx += 1
			if idx < n {
				nd.Right = nodes[idx]
			}
			if nd.Left != nil {
				lvs = append(lvs, nd.Left)
			}
			if nd.Right != nil {
				lvs = append(lvs, nd.Right)
			}
		}
		start += tmpN * 2
		lvs = lvs[tmpN:]
	}
	return root
}

func StackInOrder(root *TreeNode) (orders []int) {
	var stk []*TreeNode
	for len(stk) != 0 || root != nil {
		if root != nil {
			stk = append(stk, root)
			for root.Left != nil {
				stk = append(stk, root.Left)
				root = root.Left
			}
		}
		n := len(stk)
		root = stk[n-1]
		stk = stk[:n-1]
		orders = append(orders, root.Val)
		root = root.Right
	}
	return
}

func StackPreOrder(root *TreeNode) (orders []int) {
	if root == nil {
		return
	}
	var stk []*TreeNode
	stk = append(stk, root)
	for len(stk) != 0 {
		n := len(stk)
		root = stk[n-1]
		stk = stk[:n-1]
		orders = append(orders, root.Val)
		if root.Right != nil {
			stk = append(stk, root.Right)
		}
		if root.Left != nil {
			stk = append(stk, root.Left)
		}
	}
	return
}

func SprintNodePairs(nodes []*NodePair) string {
	strs := make([]string, len(nodes))
	for i, nd := range nodes {
		strs[i] = fmt.Sprintf("node=%+v", nd)
	}
	return strings.Join(strs, "\t")
}

type NodePair struct {
	node *TreeNode
	eof  bool
}

func StackPostOrder(root *TreeNode) (orders []int) {
	var stk []*NodePair
	for len(stk) != 0 || root != nil {
		if root != nil {
			stk = append(stk, &NodePair{root, false})
			for root.Left != nil {
				stk = append(stk, &NodePair{root.Left, false})
				root = root.Left
			}
		}
		//fmt.Printf("stack=%s\n", SprintNodePairs(stk))
		n := len(stk)
		pr := stk[n-1]
		root = pr.node
		stk = stk[:n-1]
		if pr.eof == true {
			orders = append(orders, root.Val)
			root = nil
		} else {
			pr.eof = true
			stk = append(stk, pr)
			root = root.Right
		}
	}
	return
}

func RecurInOrder(root *TreeNode) (orders []int) {
	orders = recurInOrder(root, nil)
	return orders
}

func RecurPreOrder(root *TreeNode) (orders []int) {
	orders = recurPreOrder(root, nil)
	return orders
}

func RecurPostOrder(root *TreeNode) (orders []int) {
	orders = recurPostOrder(root, nil)
	return orders
}

func recurInOrder(root *TreeNode, ods []int) (orders []int) {
	orders = ods
	if root == nil {
		return orders
	}
	orders = recurInOrder(root.Left, orders)
	orders = append(orders, root.Val)
	orders = recurInOrder(root.Right, orders)
	return
}

func recurPreOrder(root *TreeNode, ods []int) (orders []int) {
	orders = ods
	if root == nil {
		return orders
	}
	orders = append(orders, root.Val)
	orders = recurPreOrder(root.Left, orders)
	orders = recurPreOrder(root.Right, orders)
	return
}

func recurPostOrder(root *TreeNode, ods []int) (orders []int) {
	orders = ods
	if root == nil {
		return orders
	}
	orders = recurPostOrder(root.Left, orders)
	orders = recurPostOrder(root.Right, orders)
	orders = append(orders, root.Val)
	return
}

func MorrisInOrder(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
			continue
		}
		nd := root.Left
		for nd.Right != nil && nd.Right != root {
			nd = nd.Right
		}
		if nd.Right == root {
			nd.Right = nil
			res = append(res, root.Val)
			root = root.Right
			continue
		} else {
			nd.Right = root
		}
		root = root.Left
	}
	return res
}

func MorrisPreOrder(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
			continue
		}
		nd := root.Left
		for nd.Right != nil && nd.Right != root {
			nd = nd.Right
		}
		if nd.Right == root {
			nd.Right = nil
			root = root.Right
			continue
		} else {
			res = append(res, root.Val)
			nd.Right = root
		}
		root = root.Left
	}
	return res
}

func reverse(res []int) {
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
}

func MorrisPostOrder(root *TreeNode) (res []int) {
	for root != nil {
		if root.Right == nil {
			res = append(res, root.Val)
			root = root.Left
			continue
		}
		nd := root.Right
		for nd.Left != nil && nd.Left != root {
			nd = nd.Left
		}
		if nd.Left == root {
			nd.Left = nil
			root = root.Left
			continue
		} else {
			res = append(res, root.Val)
			nd.Left = root
		}
		root = root.Right
	}
	reverse(res)
	return res
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	n1 := &TreeNode{Val: 1}
	res := []*TreeNode{n1}
	for i := 2; i <= n; i++ {
		var tmp []*TreeNode
		for _, t := range res {
			ti := addNode(t, i)
			tmp = append(tmp, ti...)
		}
		res = tmp
	}
	return res
}

func addNode(t *TreeNode, x int) []*TreeNode {
	var res []*TreeNode
	if t == nil {
		nx := &TreeNode{Val: x}
		res = append(res, nx)
		return res
	}
	//Top
	nt := &TreeNode{Val: x}
	nt.Left = deepCopy(t)
	res = append(res, nt)
	//sub
	cnt := 1
	tr := t
	for tr.Right != nil {
		tr = tr.Right
		cnt++
	}
	for i := 0; i < cnt; i++ {
		nr := deepCopy(t)
		nx := &TreeNode{Val: x}
		res = append(res, nr)
		for j := 0; j < i; j++ {
			nr = nr.Right
		}
		nx.Left = nr.Right
		nr.Right = nx
	}
	return res
}

func deepCopy(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nd := &TreeNode{Val: root.Val}
	nd.Left = deepCopy(root.Left)
	nd.Right = deepCopy(root.Right)
	return nd
}
