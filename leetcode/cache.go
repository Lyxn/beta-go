package leetcode

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	kv       map[int]*Node
}

type Node struct {
	Key  int
	Val  int
	Last *Node
	Next *Node
}

func NewLRU(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Last = head
	c := LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		kv:       make(map[int]*Node, capacity),
	}
	return c
}

func (c *LRUCache) Get(key int) int {
	nd, ok := c.kv[key]
	if ok {
		c.moveToTail(nd)
		return nd.Val
	}
	return -1
}

func (c *LRUCache) moveToTail(nd *Node) {
	if nd.Next == nil {
		return
	}
	hd := nd.Last
	tl := nd.Next
	hd.Next = tl
	tl.Last = hd
	c.addToTail(nd)
}

func (c *LRUCache) addToTail(nd *Node) {
	last := c.tail.Last
	last.Next = nd
	nd.Last = last
	nd.Next = c.tail
	c.tail.Last = nd
}

func (c *LRUCache) addToTail1(nd *Node) {
	nd.Next = nil
	nd.Last = c.tail
	c.tail.Next = nd
	c.tail = nd
}

func (c *LRUCache) Put(key int, value int) {
	nd, ok := c.kv[key]
	if ok {
		nd.Val = value
		c.moveToTail(nd)
		return
	}
	size := len(c.kv)
	if size == c.capacity {
		n0 := c.head.Next
		delete(c.kv, n0.Key)
		if n0.Next != nil {
			n1 := n0.Next
			c.head.Next = n1
			n1.Last = c.head
		}
	}
	nd = &Node{Key: key, Val: value}
	c.kv[key] = nd
	c.addToTail(nd)
}
