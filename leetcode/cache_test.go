package leetcode

import "testing"

func TestNewLRU(t *testing.T) {
	var ret int
	c := NewLRU(2)
	c.Put(1, 1)
	c.Put(2, 2)
	ret = c.Get(1)
	t.Logf("ret=%v", ret)
	c.Put(3, 3)
	ret = c.Get(2)
	t.Logf("ret=%v", ret)
	c.Put(4, 4)
	ret = c.Get(1)
	t.Logf("ret=%v", ret)
	ret = c.Get(3)
	t.Logf("ret=%v", ret)
	ret = c.Get(4)
	t.Logf("ret=%v", ret)
}
