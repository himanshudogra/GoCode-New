package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3, 4, 5}
	//	slice2 := append(slice[1:])
	slice2 := slice[1:]
	fmt.Println(slice2)
}
