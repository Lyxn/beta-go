package leetcode

import "testing"

func TestTokenize(t *testing.T) {
	tests := []struct {
		code string
		want int
	}{
		{"12", 1},
		{"* xx 33", 3},
		{"(+ 1 2)", 5},
	}
	for _, tt := range tests {
		tokens := tokenize(tt.code)
		ret := len(tokens)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
		t.Logf("tokens=%v", tokens)
	}
}

func TestNextParen(t *testing.T) {
	tests := []struct {
		code string
		want int
	}{
		{"(* xx 33)", 4},
		{"(+ 1 2)", 4},
		{"(add (+ 1 2) z)", 8},
		{"(let z 5 (+ 1 2) z)", 10},
	}
	for _, tt := range tests {
		tokens := tokenize(tt.code)
		ret := nextParen(tokens, 0)
		if ret != tt.want {
			t.Errorf("get=%v want=%v", ret, tt.want)
		}
		t.Logf("tokens=%v", tokens)
	}
}

func TestNextExpr(t *testing.T) {
	tests := []struct {
		code string
		s    int
		want int
	}{
		{"(* xx 33)", 0, 4},
		{"(+ 1 2)", 2, 2},
		{"(add (+ 1 2) z)", 2, 6},
		{"(let z 5 (+ 1 2) z)", 4, 8},
	}
	for _, tt := range tests {
		tokens := tokenize(tt.code)
		ret := nextExpr(tokens, tt.s)
		if ret != tt.want {
			t.Errorf("code=%v get=%v want=%v", tt.code, ret, tt.want)
		}
	}
}

func TestEvalLispCode(t *testing.T) {
	tests := []struct {
		code string
		want int
	}{
		{"(add 12 33)", 45},
		{"(mult 1 2)", 2},
		{"(add (add 1 2) 6)", 9},
		{"(let x 3 x 2 x)", 2},
		{"(let z 5 (add (mult 1 2) z))", 7},
		{"(let x 2 (mult x 5))", 10},
		{"(let x 2 (mult x (let x 3 y 4 (add x y))))", 14},
		{"(let x 1 y 2 x (add x y) (add x y))", 5},
		{"(let x 2 (add (let x 3 (let x 4 x)) x))", 6},
		{"(let a1 3 b2 (add a1 1) b2)", 4},
		{"(let x 7 -12)", -12},
	}
	for _, tt := range tests {
		ret := evalLispCode(tt.code)
		if ret != tt.want {
			t.Errorf("code=%v get=%v want=%v", tt.code, ret, tt.want)
		}
	}
}
