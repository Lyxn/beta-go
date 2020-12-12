package leetcode

import (
	"encoding/json"
	"log"
)

const EPS = 1e-5

func DecodeJson(str string, res interface{}) {
	e := json.Unmarshal([]byte(str), res)
	if e != nil {
		log.Printf("data=%v error=%v", str, e)
	}
}

func EncodeJson(val interface{}) string {
	res, e := json.Marshal(val)
	if e != nil {
		log.Printf("data=%v error=%v", val, e)
	}
	return string(res)
}

func MinInt(a int, nums ...int) (ret int) {
	ret = a
	for _, n := range nums {
		if n < ret {
			ret = n
		}
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func isEqualFloat(a, b float64) bool {
	d := a - b
	return d >= -EPS && d < EPS
}

func isEqualInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
