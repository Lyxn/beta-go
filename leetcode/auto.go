package leetcode

type KMP struct {
	p   string
	nxt []int
}

func (k *KMP) Build0(p string) {
	n := len(p)
	nxt := make([]int, n)
	for i := 1; i < n; i++ {
		last := nxt[i-1]
		if p[last] == p[i] {
			nxt[i] = last + 1
		}
	}
	k.p = p
	nxt = append([]int{0}, nxt[:n-1]...)
	k.nxt = nxt
}

func (k *KMP) Build(p string) {
	n := len(p)
	nxt := make([]int, n)
	for i := 1; i < n-1; i++ {
		fail := nxt[i]
		if p[fail] == p[i] {
			nxt[i+1] = fail + 1
		}
	}
	for i := 2; i < n; i++ {
		if p[nxt[i]] == p[i] {
			nxt[i] = nxt[nxt[i]]
		}
	}
	k.p = p
	k.nxt = nxt
}

func (k *KMP) Search(ss string) int {
	np := len(k.p)
	ns := len(ss)
	i := 0
	j := 0
	for i < ns && j < np {
		if ss[i] == k.p[j] {
			i++
			j++
			continue
		}
		if j == 0 {
			i++
		} else {
			j = k.nxt[j]
		}
	}
	if j == np {
		return i - np
	}
	return -1
}
