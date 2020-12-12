package util

import "testing"

func TestParseInts(t *testing.T) {
	str := "[1,2,3,4]"
	buf := []byte(str)
	res := parseInts(buf)
	t.Logf("get=%v", res)
}

func TestReadInts(t *testing.T) {
	//filename := "./test.txt"
	filename := "../leetcode/case/reverse-pairs-37.txt"
	res := ReadInts(filename)
	t.Logf("get=%v", len(res[0]))
}
