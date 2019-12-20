package main

import (
	"fmt"
)

func main() {

	statepopulations := make(map[string]int)
	statepopulations = map[string]int{
		"California": 23897645,
		"Florida":    34567896,
		"Georgia":    43566433,
	}
	fmt.Println(statepopulations)
	statepopulations["Texas"] = 45867922
	fmt.Println(statepopulations)
}
