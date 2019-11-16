package main

import (
	"fmt"
)

const (
	a = iota // iota is an int type data.
	b        // If you don't assign any value then the compiler will look for the pattern and continue it with the const scope.
	c        // If you don't assign any value then the compiler will look for the pattern and continue it with the const scope.
)

const (
	a2 = iota
)

func main() {

	fmt.Printf("%v,%T\n", a, a)
	fmt.Printf("%v,%T\n", b, b)
	fmt.Printf("%v,%T\n", c, c)
	fmt.Printf("%v,%T\n", a2, a2)
}
