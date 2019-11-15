package main

import (
	"fmt"
)

func main() {

	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // Exclusive OR (means, one has the bit set and as another, but not both) 1001
	fmt.Println(a &^ b) // AND NOT, (it's gonna be set to true of neighter one has the bit set) 0100
}
