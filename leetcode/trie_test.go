package leetcode

import (
	"testing"
)

func TestTrie0(t *testing.T) {
	trie := NewTrie()
	var ret bool
	trie.Insert("apple")
	ret = trie.Search("apple")
	t.Logf("1 get=%v", ret)
	ret = trie.StartsWith("app")
	t.Logf("2 get=%v", ret)
}

func TestWordBreak(t *testing.T) {
	//ss := "pineapplepenapple"
	//dct := []string{"pine", "apple", "pineapple", "pen"}
	//ss := "catsanddog"
	//dct := []string{"cat", "cats", "and", "sand", "dog"}
	ss := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	//ss := "aaaaaaaaaaaaaaaaaaaaa"
	dct := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}

	ret := wordBreak2(ss, dct)
	t.Logf("get=%v", len(ret))
}
