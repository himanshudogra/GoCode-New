package main

import (
	"fmt"
)

func main() {

	statepopulations := make(map[string]int)

	statepopulations = map[string]int{
		"Rajasthan": 348372920,
		"UP":        44759388992,
		"MP":        3334940093,
	}
	fmt.Println(len(statepopulations))
}
