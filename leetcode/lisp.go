package leetcode

import (
	"math"
)

const InvalidCode = math.MaxInt32

func isWhite(c byte) bool {
	return c == ' ' || c == '\n'
}

func isParen(c byte) bool {
	return c == '(' || c == ')'
}

func tokenize(code string) (res []string) {
	n := len(code)
	i := 0
	for i < n {
		for i < n && isWhite(code[i]) {
			i++
		}
		if i >= n {
			break
		}
		if isParen(code[i]) {
			res = append(res, code[i:i+1])
			i++
		} else {
			s := i
			for i < n && (!isWhite(code[i]) && !isParen(code[i])) {
				i++
			}
			res = append(res, code[s:i])
		}
	}
	return
}

func nextParen(tokens []string, s int) int {
	n := len(tokens)
	if s >= n || tokens[s] != "(" {
		return -1
	}
	cnt := 1
	i := s + 1
	for ; i < n; i++ {
		if tokens[i] == "(" {
			cnt++
		} else if tokens[i] == ")" {
			cnt--
		}
		if cnt == 0 {
			break
		}
	}
	if i >= n {
		return -1
	}
	return i
}

func nextExpr(tokens []string, s int) (e int) {
	if tokens[s] == "(" {
		return nextParen(tokens, s)
	} else {
		return s
	}
}

func isVar(ss string) bool {
	if isStrDigit(ss) {
		return false
	} else if isParen(ss[0]) {
		return false
	}
	return true
}

func isStrDigit(ss string) bool {
	n := len(ss)
	i := 0
	if ss[i] == '-' {
		i++
	}
	for ; i < n; i++ {
		if ss[i] < '0' || ss[i] > '9' {
			return false
		}
	}
	return true
}

func parseInt(ss string) int {
	if ss[0] == '-' {
		return -parseInt(ss[1:])
	}
	ret := 0
	for _, c := range ss {
		ret = ret*10 + int(c-'0')
	}
	return ret
}

func findEnv(env []map[string]int, key string) int {
	n := len(env)
	for i := n - 1; i >= 0; i-- {
		val, ok := env[i][key]
		if ok {
			return val
		}
	}
	return InvalidCode
}

func evalLispByTokens(env []map[string]int, tokens []string, s, e int) (ret int) {
	if s > e {
		return InvalidCode
	}
	if tokens[s] != "(" {
		expr := tokens[s]
		if isStrDigit(expr) {
			return parseInt(expr)
		} else {
			return findEnv(env, expr)
		}
	}
	i := s + 1
	if tokens[i] == "add" {
		s1 := i + 1
		e1 := nextExpr(tokens, s1)
		s2 := e1 + 1
		e2 := nextExpr(tokens, s2)
		return evalLispByTokens(env, tokens, s1, e1) + evalLispByTokens(env, tokens, s2, e2)
	} else if tokens[i] == "mult" {
		s1 := i + 1
		e1 := nextExpr(tokens, s1)
		s2 := e1 + 1
		e2 := nextExpr(tokens, s2)
		return evalLispByTokens(env, tokens, s1, e1) * evalLispByTokens(env, tokens, s2, e2)
	}
	if tokens[i] != "let" {
		return InvalidCode
	}
	kv := make(map[string]int)
	env = append(env, kv)
	k := i + 1
	for k+1 < e && isVar(tokens[k]) {
		vs := k + 1
		ve := nextExpr(tokens, vs)
		kv[tokens[k]] = evalLispByTokens(env, tokens, vs, ve)
		k = ve + 1
	}
	s1 := k
	e1 := nextExpr(tokens, s1)
	ret = evalLispByTokens(env, tokens, s1, e1)
	env = env[:len(env)-1]
	return ret
}

func evalLispCode(code string) int {
	tokens := tokenize(code)
	env := make([]map[string]int, 0)
	return evalLispByTokens(env, tokens, 0, len(tokens)-1)
}
