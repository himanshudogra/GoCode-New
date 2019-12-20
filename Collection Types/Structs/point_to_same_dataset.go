package main

import (
	"fmt"
)

func main() {

	aDoctor := struct{ name string }{name: "Tom Berry"} // aDoctor is a struct itself.
	anotherDoctor := &aDoctor                           // anotherDoctor is a pointer to that struct.
	anotherDoctor.name = "Yo Yo"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}
