package leetcode

import "strings"

type Trie struct {
	Eof      bool
	Children [26]*Trie
}

/** Initialize your data structure here. */
func NewTrie() Trie {
	trie := Trie{
		Eof:      false,
		Children: [26]*Trie{},
	}
	return trie
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		i := c - 'a'
		nd := node.Children[i]
		if nd == nil {
			n := NewTrie()
			node.Children[i] = &n
		}
		node = node.Children[i]
	}
	node.Eof = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t
	for _, c := range word {
		i := c - 'a'
		node = node.Children[i]
		if node == nil {
			return false
		}
	}
	return node.Eof
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, c := range prefix {
		i := c - 'a'
		node = node.Children[i]
		if node == nil {
			return false
		}
	}
	return true
}

type pairTrie struct {
	trie *Trie
	idx  int
}

func (t *Trie) GetAllPrefix(ss string, si int) []int {
	var res []int
	ns := len(ss)
	if si >= ns {
		return res
	}
	stk := []pairTrie{{t, si}}
	for len(stk) > 0 {
		n := len(stk)
		pr := stk[n-1]
		stk = stk[:n-1]
		i := ss[pr.idx] - 'a'
		node := pr.trie.Children[i]
		if node != nil {
			if node.Eof {
				res = append(res, pr.idx)
			}
			if pr.idx+1 < ns {
				stk = append(stk, pairTrie{node, pr.idx + 1})
			}
		}
	}
	return res
}

func wordBreak0(ss string, wordDict []string) []string {
	trie := NewTrie()
	for _, word := range wordDict {
		trie.Insert(word)
	}
	ns := len(ss)
	stk := [][]int{{0}}
	res := [][]int{}
	for len(stk) > 0 {
		ids := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		ps := trie.GetAllPrefix(ss, ids[len(ids)-1])
		if len(ps) == 0 {
			continue
		}
		for _, ei := range ps {
			tmp := append([]int{}, ids...)
			tmp = append(tmp, ei+1)
			if ei+1 == ns {
				res = append(res, tmp)
			}
			stk = append(stk, tmp)
		}
	}
	resStr := make([]string, len(res))
	for i, ids := range res {
		resStr[i] = getStrByIdx(ss, ids)
	}
	return resStr
}

func wordBreak1(ss string, wordDict []string) []string {
	trie := NewTrie()
	for _, word := range wordDict {
		trie.Insert(word)
	}
	ns := len(ss)
	dp := make([][][]string, ns+1)
	dp[ns] = [][]string{{}}
	var dfs func(int) [][]string
	dfs = func(i int) [][]string {
		if dp[i] != nil {
			return dp[i]
		}
		ps := trie.GetAllPrefix(ss, i)
		res := make([][]string, 0)
		for _, e := range ps {
			word := ss[i : e+1]
			for _, words := range dfs(e + 1) {
				next := append([]string{word}, words...)
				res = append(res, next)
			}
		}
		dp[i] = res
		return res
	}

	res := make([]string, 0)
	for _, r := range dfs(0) {
		res = append(res, strings.Join(r, " "))
	}
	return res
}

func wordBreak2(ss string, wordDict []string) []string {
	wordSet := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}
	ns := len(ss)
	dp := make([][][]string, ns+1)
	dp[ns] = [][]string{{}}
	var dfs func(int) [][]string
	dfs = func(i int) [][]string {
		if dp[i] != nil {
			return dp[i]
		}
		res := make([][]string, 0)
		for e := i + 1; e <= ns; e++ {
			word := ss[i:e]
			if _, ok := wordSet[word]; !ok {
				continue
			}
			for _, words := range dfs(e) {
				next := append([]string{word}, words...)
				res = append(res, next)
			}
		}
		dp[i] = res
		return res
	}

	res := make([]string, 0)
	for _, r := range dfs(0) {
		res = append(res, strings.Join(r, " "))
	}
	return res
}

func wordBreak(ss string, wordDict []string) []string {
	trie := NewTrie()
	for _, word := range wordDict {
		trie.Insert(word)
	}
	ns := len(ss)
	dp := make([][][]string, ns+1)
	dp[ns] = [][]string{{}}
	stk := [][2]int{{0, 0}}
	for len(stk) > 0 {
		cp := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		s := cp[0]
		if dp[s] != nil {
			continue
		}
		ps := trie.GetAllPrefix(ss, s)
		if cp[1] == 0 {
			stk = append(stk, [2]int{s, 1})
			for _, e := range ps {
				stk = append(stk, [2]int{e + 1, 0})
			}
			continue
		}
		for _, e := range ps {
			word := ss[s : e+1]
			for _, r := range dp[e+1] {
				rs := append([]string{word}, r...)
				dp[s] = append(dp[s], rs)
			}
		}
	}

	res := make([]string, len(dp[0]))
	for i, r := range dp[0] {
		res[i] = strings.Join(r, " ")
	}
	return res
}

func getStrByIdx(ss string, ids []int) string {
	res := make([]string, len(ids)-1)
	for i := 0; i < len(ids)-1; i++ {
		res[i] = ss[ids[i]:ids[i+1]]
	}
	return strings.Join(res, " ")
}
