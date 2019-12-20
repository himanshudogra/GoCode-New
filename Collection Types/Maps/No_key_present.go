package main

import (
	"fmt"
)

func main() {

	statepopulations := make(map[string]int)
	statepopulations = map[string]int{
		"California": 5434566,
		"Texas":      54384759,
	}

	fmt.Println(statepopulations["calfornia"]) // this means that the key is not found in the program.
}
