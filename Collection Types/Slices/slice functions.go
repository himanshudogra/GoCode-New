package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3}
	b := a
	b[1] = 4

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

}
