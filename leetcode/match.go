package leetcode

func isAllStar(p string) bool {
	for _, c := range p {
		if c != '*' {
			return false
		}
	}
	return true
}

func isMatchWildGreedy(s string, p string) bool {
	ns := len(s)
	np := len(p)
	si := 0
	pi := 0
	for si < ns && pi < np {
		if p[pi] == '*' {
			break
		} else if p[pi] != '?' && p[pi] != s[si] {
			return false
		}
		si++
		pi++
	}
	for ns > si && np > pi {
		if p[np-1] == '*' {
			break
		} else if p[np-1] != '?' && p[np-1] != s[ns-1] {
			return false
		}
		ns--
		np--
	}
	s = s[:ns]
	p = p[:np]
	for si < ns && pi < np {
		spi := pi
		for spi < np && p[spi] == '*' {
			spi++
		}
		if spi == np {
			return true
		}
		epi := spi
		for epi < np && p[epi] != '*' {
			epi++
		}
		si = findStrWild(s, p[spi:epi], si)
		if si == -1 {
			return false
		}
		pi = epi
	}
	if si == ns && pi == np {
		return true
	} else if pi == np {
		return false
	}
	return isAllStar(p[pi:])
}

func findStrWild(s, p string, si int) int {
	l := len(p)
	ns := len(s)
	for i := si; i <= ns-l; i++ {
		if isEqualWild(s[i:i+l], p) {
			return i + l
		}
	}
	return -1
}

func isEqualWild(s, p string) bool {
	n := len(s)
	for i := 0; i < n; i++ {
		if p[i] != '?' && s[i] != p[i] {
			return false
		}
	}
	return true
}
