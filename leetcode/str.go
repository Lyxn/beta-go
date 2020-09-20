package leetcode

func isValidParen(sr string) bool {
	var stk []rune
	for i := 0; i < len(sr); i++ {
		if sr[i] == '(' {
			stk = append(stk, ')')
		} else if sr[i] == '{' {
			stk = append(stk, '}')
		} else if sr[i] == '[' {
			stk = append(stk, ']')
		} else if sr[i] == ')' || sr[i] == '}' || sr[i] == ']' {
			num := len(stk)
			if num == 0 {
				return false
			} else if stk[num-1] != rune(sr[i]) {
				return false
			}
			stk = stk[:num-1]
		}
	}
	return len(stk) == 0
}
