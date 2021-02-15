package leetcode

//UnionFind Int
type UFInt struct {
	key    int
	depth  int
	parent *UFInt
}

func (u *UFInt) Find() *UFInt {
	c := u
	for c.parent != nil {
		c = c.parent
	}
	return c
}

func (u *UFInt) Union(other *UFInt) {
	lhs := u.Find()
	rhs := other.Find()
	if lhs == rhs {
		return
	}
	if lhs.depth == rhs.depth {
		rhs.parent = lhs
		lhs.depth += 1
	} else if lhs.depth > rhs.depth {
		rhs.parent = lhs
	} else {
		lhs.parent = rhs
	}
}

func CountUFInt(nodes []*UFInt) int {
	res := 0
	for _, nd := range nodes {
		if nd.parent == nil {
			res += 1
		}
	}
	return res
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	nodes := make([]*UFInt, n)
	for i := 0; i < n; i++ {
		nodes[i] = &UFInt{key: i}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || isConnected[i][j] == 0 {
				continue
			}
			nodes[i].Union(nodes[j])
		}
	}
	return CountUFInt(nodes)
}

//UnionFind String
type UFStr struct {
	key    string
	val    float64
	depth  int
	parent *UFStr
}

func (u *UFStr) FindVal() (*UFStr, float64) {
	c := u
	v := c.val
	for c.parent != nil {
		c = c.parent
		v *= c.val
	}
	return c, v
}

func UnionVal(a, b *UFStr, ab float64) {
	c, ac := a.FindVal()
	d, bd := b.FindVal()
	if c == d {
		return
	}
	if c.depth == d.depth {
		c.parent = d
		c.val = ab * bd / ac
		c.depth += 1
	} else if c.depth > d.depth {
		d.parent = c
		d.val = ac / (ab * bd)
	} else {
		c.parent = d
		c.val = ab * bd / ac
	}
}

func buildUFNodes(equations [][]string, values []float64) (nodes map[string]*UFStr) {
	nodes = make(map[string]*UFStr)
	for i, pr := range equations {
		a := pr[0]
		b := pr[1]
		ab := values[i]
		nda, oka := nodes[a]
		if !oka {
			nda = &UFStr{key: a, val: 1}
			nodes[a] = nda
		}
		ndb, okb := nodes[b]
		if !okb {
			ndb = &UFStr{key: b, val: 1}
			nodes[b] = ndb
		}
		UnionVal(nda, ndb, ab)
	}
	return nodes
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	nodes := buildUFNodes(equations, values)

	dfs := func(a, b string) float64 {
		nda, oka := nodes[a]
		ndb, okb := nodes[b]
		if !oka || !okb {
			return -1
		} else if a == b {
			return 1
		}
		ra, ar := nda.FindVal()
		rb, br := ndb.FindVal()
		if ra != rb {
			return -1
		}
		return ar / br
	}

	res := make([]float64, len(queries))
	for i, pr := range queries {
		res[i] = dfs(pr[0], pr[1])
	}
	return res
}
