package main

import (
	"fmt"
)

func main() {

	var i int = 43
	fmt.Printf("%v,%T\n", i, i)

	var j float32
	j = float32(i) // Here float32 dataType is being used as a func.
	fmt.Printf("%v,%T\n", j, j)
}
