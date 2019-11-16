package main

import (
	"fmt"
)

func main() {

	var n string = "This is a string"
	fmt.Printf("%v,%T\n", string(n[2]), string(n[2]))
}
