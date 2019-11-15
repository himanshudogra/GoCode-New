package main

import (
	"fmt"
)

var i int = 23

func main() {
	var i int = 43 // this will take the precedence and hide the package level variable.
	fmt.Println(i)

}
