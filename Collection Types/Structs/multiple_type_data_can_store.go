package main

import (
	"fmt"
)

type Doctor struct { // struct declaration
	Number     int
	Name       string
	Companions []string
}

func main() {

	aDoctor := Doctor{ // defining the different type of values under the struct.
		Number:     34,
		Name:       "Lisa Roy",
		Companions: []string{"Jack Mars", "stephen jones", "pee writo"},
	}
	fmt.Println(aDoctor)
}
