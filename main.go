package main

import "fmt"

func main() {
	c := []int{1,2,3}
	c[1], c[0] = c[0], c[1]
	fmt.Printf("hello, %v", c)
}
