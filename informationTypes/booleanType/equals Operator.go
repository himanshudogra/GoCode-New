package main

import (
	"fmt"
)

func main() {

	var i = 1 == 1
	var j = 1 == 2
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", j, j)
}
