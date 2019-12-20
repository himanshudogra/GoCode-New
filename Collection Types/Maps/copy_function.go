package main

import (
	"fmt"
)

func main() {

	statepopulations := make(map[string]int)
	statepopulations = map[string]int{

		"Rajasthan":   239474003,
		"UP":          384783893,
		"MP":          392889923,
		"Maharashtra": 28389923839,
		"Punjab":      23829392323,
	}
	new := statepopulations // both are pointing to the same underlying data.
	delete(new, "Rajasthan")
	fmt.Println(new)
	fmt.Println(statepopulations)

}
