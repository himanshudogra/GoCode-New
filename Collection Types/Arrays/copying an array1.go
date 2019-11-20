package main

import (
	"fmt"
)

func main() {

	a := [...]int{1, 3, 4}
	b := &a
	b[1] = 4

	fmt.Println(a)
	fmt.Println(b)
}
