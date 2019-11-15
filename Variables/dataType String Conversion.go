package main

import (
	"fmt"
	"strconv" // this package supports the string conversion.
)

func main() {

	var i int = 32
	fmt.Printf("%v,%T\n", i, i)

	var j string
	j = strconv.Itoa(i) // this strconv function is used to convert integer type to ascii type.
	fmt.Printf("%v,%T\n", j, j)
}
