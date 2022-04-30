package leetcode

import (
	"fmt"
	"testing"
)

func TestBuildBinaryTree(t *testing.T) {
	nums := []int{1, 2, 3, NilNode, NilNode, 5, 6}
	root := BuildBinaryTree(nums)
	t.Logf("root=%+v", root)
	var orders []int
	//orders = StackInOrder(root)
	//t.Logf("InOrders=%v", orders)
	//orders = StackPreOrder(root)
	//t.Logf("PreOrders=%v", orders)
	//orders = StackPostOrder(root)
	//t.Logf("PostOrders=%v", orders)
	//orders = RecurInOrder(root)
	//t.Logf("InOrders=%v", orders)
	//orders = RecurPreOrder(root)
	//t.Logf("PreOrders=%v", orders)
	orders = RecurPostOrder(root)
	t.Logf("PostOrders=%v", orders)
	//orders = MorrisInOrder(root)
	//t.Logf("InOrders=%v", orders)
	orders = MorrisPreOrder(root)
	t.Logf("PreOrders=%v", orders)
	orders = MorrisPostOrder(root)
	t.Logf("PostOrders=%v", orders)
}

func TestSerialize(t *testing.T) {
	//str := "[1,2,3,null,null,4,5,6,7]"
	str := "[1,2,3,null,8,4,5,6,7,null,null,9,null,10,11,12,13,14,15,16]"
	root := deserialize(str)
	ret := serialize(root)
	t.Logf("tree=%v", str)
	if str != ret {
		t.Errorf("want=%s get=%s", str, ret)
	}
}

func TestGenerateTrees(t *testing.T) {
	res := generateTrees(3)
	for _, r := range res {
		fmt.Println(serialize(r))
	}
}
