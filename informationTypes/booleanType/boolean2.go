package main

import (
	"fmt"
)

func main() {

	var a bool
	// here we have not initialize any value under variable a,
	// so bydefault in GO, boolean variable assign the value Zero, which is false.
	fmt.Printf("%v,%T\n", a, a)

}
