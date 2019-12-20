package main

import (
	"fmt"
)

func main() {

	statepopulations := make(map[string]int)
	statepopulations = map[string]int{

		"california": 4558373,
		"Florida":    34857458,
	}
	_, ok := statepopulations["california"] // using writeonly variable to throw away the value to int
	fmt.Println(ok)
}
