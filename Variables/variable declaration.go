package main

import (
	"fmt"
)

func main() {
	var i int = 43
	var j float32
	j = 54.9
	k := 32

	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", j, j)
	fmt.Printf("%v,%T\n", k, k)
}
