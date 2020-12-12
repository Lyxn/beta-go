package leetcode

import "testing"

func TestLastStoneWeight(t *testing.T) {
	//stones := []int{1,2,1}
	stones := []int{2, 7, 4, 1, 8, 1}
	ret := lastStoneWeightII(stones)
	t.Logf("ret=%v", ret)
}

func TestProfitableSchema(t *testing.T) {
	tests := []struct {
		want   int
		G      int
		P      int
		group  []int
		profit []int
	}{
		{
			4,
			2,
			1,
			[]int{1, 1, 2},
			[]int{1, 2, 3},
		},
		{
			2,
			5,
			3,
			[]int{2, 2},
			[]int{2, 3},
		},
	}
	for _, tt := range tests {
		ret := profitableSchemes1(tt.G, tt.P, tt.group, tt.profit)
		if ret != tt.want {
			t.Errorf("want=%v get=%v", tt.want, ret)
		}
	}
}

func TestWordsTyping(t *testing.T) {
	//sents := []string{"a"}
	//rows := 20000
	//cols := 20000
	//sents := []string{"hello", "world"}
	//rows := 10
	//cols := 11
	sents := []string{"hello", "man", "girl", "boy", "dog", "cat", "giggle", "mister", "risk", "code", "knight", "slack", "telegram", "china", "usa", "japan", "korea", "india", "russia", "australia", "may", "be", "like", "you", "love", "preview", "photo", "picture", "hello", "man", "girl", "boy", "dog", "cat", "giggle", "mister", "risk", "code", "knight", "slack", "telegram", "china", "usa", "japan", "korea", "india", "russia", "australia", "may", "be", "like", "you", "love", "preview", "photo", "picture"}
	rows := 1
	cols := 10000

	ret := wordsTyping(sents, rows, cols)
	t.Logf("ret=%v", ret)
}
