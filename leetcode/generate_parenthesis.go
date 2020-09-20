package leetcode

const (
	LPar = "("
	RPar = ")"
)

func generateParenthesis(n int) (res []string) {
	res = dfsGenPar(res, LPar, 1, n-1)
	return res
}

func dfsGenPar(strs []string, str string, open, remain int) (res []string) {
	res = strs
	if remain == 0 {
		for i := 0; i < open; i++ {
			str += LPar
		}
		res = append(res, str)
		return res
	}
	str1 := str + LPar
	res = dfsGenPar(res, str1, open+1, remain-1)
	if open > 0 {
		str2 := str + RPar
		res = dfsGenPar(res, str2, open-1, remain)
	}
	return res
}
