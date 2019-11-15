package main

import (
	"fmt"
)

func main() {

	a := 8 // 2^3

	fmt.Println(a << 2) // 2^3 * 2^2 = 2^5 = 32
	fmt.Println(a >> 2) // 2^3 / 2^2 = 2^1 = 2
}

// Note: learn more about bit shifting: https://golang.org/ref/spec#Iota and https://stackoverflow.com/questions/5801008/go-and-operators
