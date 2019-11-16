package main

import (
	"fmt"
)

func main() {

	const a = 42
	var b int16 = 43
	fmt.Printf("%v,%T\n", a+b, a+b) // In this case compiler is interpreting the a as int16. ex: 42 + b it takes.
}
