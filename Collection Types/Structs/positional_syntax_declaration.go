package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	Name       string
	companions []string
}

func main() {

	aDoctor := Doctor{
		23,
		"Himanshu Dogra",
		[]string{"Lisa Roy"},
	} // The problem with this is if you add/remove any further values from the struct then
	// it will change the oder to pick up the right value.
	fmt.Println(aDoctor.companions)
}
