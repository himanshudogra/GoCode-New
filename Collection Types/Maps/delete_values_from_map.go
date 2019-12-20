package main

import (
	"fmt"
)

func main() {

	statepopulations := make(map[string]int)
	statepopulations = map[string]int{
		"California": 3457892,
		"Florida":    463782,
		"Georgia":    8755938,
	}
	fmt.Println(statepopulations)
	statepopulations["Texas"] = 456783
	fmt.Println(statepopulations)
	delete(statepopulations, "Texas")
	fmt.Println(statepopulations)
}
