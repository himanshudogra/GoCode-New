package main

import (
	"fmt"
)

func main() {

	statepopulations := make(map[string]int)
	statepopulations = map[string]int{
		"California": 23897645,
	}
	fmt.Println(statepopulations["California"])
}
