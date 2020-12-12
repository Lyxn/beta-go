package leetcode

import (
	"strconv"
	"strings"
)

func hanota(A, B, C []int) []int {
	A, B, C = moveHanota(len(A), A, B, C)
	return C
}

func moveHanota(n int, A []int, B []int, C []int) (A1, B1, C1 []int) {
	if n == 1 {
		C = append(C, A[len(A)-1])
		A = A[:len(A)-1]
	} else {
		A, C, B = moveHanota(n-1, A, C, B)
		C = append(C, A[len(A)-1])
		A = A[:len(A)-1]
		B, A, C = moveHanota(n-1, B, A, C)
	}
	return A, B, C
}

func decodeString(ss string) string {
	n := len(ss)
	var stk []string
	s := []rune(ss)
	for i := 0; i < n; {
		if isLetter(s[i]) {
			si := i
			for i < n && isLetter(s[i]) {
				i++
			}
			stk = append(stk, string(s[si:i]))
			continue
		} else if isDigit(s[i]) {
			si := i
			for i < n && isDigit(s[i]) {
				i++
			}
			stk = append(stk, string(s[si:i]))
			continue
		} else if s[i] == '[' {
			i++
			continue
		}
		str := ""
		for len(stk) > 0 && !isNum(stk[len(stk)-1]) {
			str = stk[len(stk)-1] + str
			stk = stk[:len(stk)-1]
		}
		ks := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		num, e := strconv.Atoi(ks)
		if e != nil {
			return ""
		}
		stk = append(stk, mutStr(str, num))
		i++
	}
	return strings.Join(stk, "")
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isNum(s string) bool {
	return isDigit(rune(s[0]))
}

func mutStr(s string, k int) string {
	ss := ""
	for i := 0; i < k; i++ {
		ss += s
	}
	return ss
}

func removeKdigits(num string, k int) string {
	n := len(num)
	stk := make([]byte, 0, n)
	i := 0
	for ; i < n; i++ {
		if len(stk) == 0 {
			stk = append(stk, num[i])
			continue
		}
		for k > 0 && len(stk) > 0 && stk[len(stk)-1] > num[i] {
			stk = stk[:len(stk)-1]
			k--
		}
		stk = append(stk, num[i])
		if k == 0 {
			break
		}
	}
	res := string(stk)
	if i != n {
		res += num[i+1:]
	}
	res = res[:len(res)-k]
	j := 0
	for ; j < len(res); j++ {
		if res[j] != '0' {
			break
		}
	}
	res = res[j:]
	if res == "" {
		return "0"
	}
	return res
}

func removeDuplicateLetters(s string) string {
	hs := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		hs[s[i]] = i
	}
	stk := make([]byte, 0)
	vis := make(map[byte]bool)
	for i := 0; i < n; i++ {
		if vis[s[i]] == true {
			continue
		}
		for len(stk) > 0 {
			top := stk[len(stk)-1]
			if top >= s[i] && hs[top] > i {
				stk = stk[:len(stk)-1]
				vis[top] = false
			} else {
				break
			}
		}
		stk = append(stk, s[i])
		vis[s[i]] = true
	}
	return string(stk)
}
