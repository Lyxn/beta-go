package util

import (
	"bufio"
	"io"
	"log"
	"os"
)

func parseInts(buf []byte) (res []int) {
	if buf[0] != '[' {
		return
	}
	n := len(buf)
	i := 1
	for i < n {
		for i < n && (buf[i] < '0' || buf[i] > '9') {
			i++
		}
		if i == n {
			break
		}
		a := 0
		for i < n && (buf[i] >= '0' && buf[i] <= '9') {
			a = a*10 + int(buf[i]-'0')
			i++
		}
		res = append(res, a)
	}
	return
}

func ReadInts(filename string) (res [][]int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("fail to read from file=%v err=%v", filename, err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			if len(buf) >= 2 {
				nums := parseInts(buf)
				res = append(res, nums)
			}
			break
		} else if err != nil {
			log.Fatalf("read line err=%v", err)
		}
		nums := parseInts(buf)
		res = append(res, nums)
	}
	return res
}
