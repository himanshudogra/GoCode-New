package main

import (
	"fmt"
)

func main() {

	var a complex64 = 34 + 67i
	fmt.Printf("%v,%T\n", real(a), real(a))
	fmt.Printf("%v,%T\n", imag(a), imag(a))
}
