package main

import (
	"fmt"
)

func main() {

	statepopulations := make(map[string]int)

	statepopulations = map[string]int{

		"California": 4568363,
		"Texas":      5363892,
	}
	pop, ok := statepopulations["California"] // (,OK) syntax will tell you if the key is present in the program or not.
	fmt.Println(pop, ok)
}
