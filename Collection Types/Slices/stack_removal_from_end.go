package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3, 4, 5}
	//	slice2 := append(slice[:4]) -> This also will work
	// slice2 := slice[:4] -> This also will work
	slice2 := slice[:len(slice)-1]
	fmt.Println(slice2)
}
