package main

import (
	"fmt"
)

func main() {

	statepopulations := make(map[string]int)
	statepopulations = map[string]int{
		"Karnataka": 34533234,
		"Rajasthan": 34568930,
	}
	fmt.Println(statepopulations)
}
