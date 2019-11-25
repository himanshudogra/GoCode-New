package main

import (
	"fmt"
)

func main() {

	a := make([]int, 3)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("capacity: %v\n", cap(a))
}
