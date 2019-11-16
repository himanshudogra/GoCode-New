package main

import (
	"fmt"
)

func main() {

	a := 16 // 2^3

	fmt.Println(a << 4) // 16 * (2)^4
	fmt.Println(a >> 4) // 16 / (2)^4
}

// Note: learn more about bit shifting: https://golang.org/ref/spec#Iota and https://stackoverflow.com/questions/5801008/go-and-operators
