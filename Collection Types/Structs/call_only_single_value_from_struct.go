package main

import (
	"fmt"
)

type Doctor struct {
	Number     int
	Name       string
	Companions []string
}

func main() {
	aDoctor := Doctor{
		Number:     23,
		Name:       "Himanshu Dogra",
		Companions: []string{"Arvind", "Prasad", "Jack"},
	}
	fmt.Println(aDoctor.Name) // That's how you can call a single value out of a complete struct.
}
