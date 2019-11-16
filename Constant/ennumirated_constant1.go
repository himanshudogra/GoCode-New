package main

import (
	"fmt"
)

const (
	_ = iota + 5
	a // iota + 5 = 1+5 = 6
	b // iota +5 = 2+5 = 7
	c // iota +5 = 3+5=8
)

func main() {

	var d int = 6
	fmt.Printf("%v\n", a == d)
}
