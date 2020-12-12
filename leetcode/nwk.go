package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	arcs := buildArcs(wordList)
	if _, ok := arcs[endWord]; !ok {
		return 0
	}
	n := len(wordList)
	db := make(map[string]int, n)
	db[beginWord] = 1
	qb := arcs[beginWord]
	cnt := 1
	for len(qb) > 0 {
		nbs := make([]string, 0)
		for _, nb := range qb {
			if nb == endWord {
				return cnt + 1
			}
			if _, okb := db[nb]; okb {
				continue
			}
			db[nb] = cnt + 1
			nb2 := arcs[nb]
			for _, w := range nb2 {
				if _, okw := db[w]; !okw {
					nbs = append(nbs, w)
				}
			}
		}
		qb = nbs
		cnt++
	}
	return 0
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	arcs := buildArcs(wordList)
	if _, ok := arcs[endWord]; !ok {
		return 0
	}
	n := len(wordList)
	db := make(map[string]int, n)
	db[beginWord] = 1
	qb := arcs[beginWord]
	de := make(map[string]int, n)
	de[endWord] = 1
	qe := arcs[endWord]
	cntBeg := 1
	cntEnd := 1
	for len(qb) > 0 && len(qe) > 0 {
		nbs := make([]string, 0)
		for _, nb := range qb {
			if c, oke := de[nb]; oke {
				return c + cntBeg
			}
			if _, okb := db[nb]; okb {
				continue
			}
			db[nb] = cntBeg + 1
			nb2 := arcs[nb]
			for _, w := range nb2 {
				if _, okw := db[w]; !okw {
					nbs = append(nbs, w)
				}
			}
		}
		qb = nbs
		cntBeg++

		nes := make([]string, 0)
		for _, ne := range qe {
			if c, okb := db[ne]; okb {
				return c + cntEnd
			}
			if _, oke := de[ne]; oke {
				continue
			}
			de[ne] = cntEnd + 1
			ne2 := arcs[ne]
			for _, w := range ne2 {
				if _, okw := de[w]; !okw {
					nes = append(nes, w)
				}
			}
		}
		qe = nes
		cntEnd++
	}
	return 0
}

func buildArcs(wordList []string) map[string][]string {
	n := len(wordList)
	arcs := make(map[string][]string, n)
	for i := 0; i < n; i++ {
		w1 := wordList[i]
		for j := i + 1; j < n; j++ {
			w2 := wordList[j]
			if !isArc(w1, w2) {
				continue
			}
			arcs[w1] = append(arcs[w1], w2)
			arcs[w2] = append(arcs[w2], w1)
		}
	}
	return arcs
}

func isArc(a, b string) bool {
	n := len(a)
	cnt := 0
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt == 1
}

type Nwk map[byte]map[byte]struct{}

func addArc(nwk Nwk, s, d byte) {
	if s == d {
		return
	}
	if _, ok := nwk[s]; !ok {
		nwk[s] = make(map[byte]struct{}, 26)
	}
	nwk[s][d] = struct{}{}
}

func getTopo(nwk Nwk) []byte {
	res := []byte{}
	for len(nwk) > 0 {
		nd, ok := getTop(nwk)
		if !ok {
			return nil
		}
		res = append(res, nd)
		delete(nwk, nd)
	}
	return res
}

func getTop(nwk Nwk) (nd byte, has bool) {
	nds := map[byte]struct{}{}
	for n := range nwk {
		nds[n] = struct{}{}
	}
	for _, ns := range nwk {
		for n := range ns {
			_, ok := nds[n]
			if ok {
				delete(nds, n)
			}
		}
	}
	if len(nds) == 0 {
		return 0, false
	}
	for nd = range nds {
		break
	}
	return nd, true
}

func alienOrder(words []string) string {
	n := len(words)
	if n == 0 {
		return ""
	}
	nwk := make(Nwk, 26)
	w0 := words[0]
	for i := 1; i < n; i++ {
		w1 := words[i]
		j := 0
		for ; j < len(w1) && j < len(w0); j++ {
			if w0[j] != w1[j] {
				addArc(nwk, w0[j], w1[j])
				break
			}
		}
		if j == len(w1) && j < len(w0) {
			return ""
		}
		w0 = w1
	}
	nds := getNodes(words)
	res := []byte{}
	if len(nwk) == 0 {
		for n := range nds {
			res = append(res, n)
		}
		return string(res)
	}
	topo := getTopo(nwk)
	if topo == nil {
		return ""
	}
	for _, n := range topo {
		delete(nds, n)
	}
	for n := range nds {
		topo = append(topo, n)
	}
	return string(topo)
}

func getNodes(words []string) map[byte]struct{} {
	nds := make(map[byte]struct{})
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			nds[w[i]] = struct{}{}
		}
	}
	return nds
}
