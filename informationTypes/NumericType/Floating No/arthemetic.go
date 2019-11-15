package main

import (
	"fmt"
)

func main() {

	var a float64 = 13.74

	a = 13.45e23
	a = 2.4E22
	fmt.Printf("%v,%T\n", a, a)
}
