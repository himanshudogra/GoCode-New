package main

import (
	"fmt"
)

func main() {

	aDoctor := struct{ name string }{name: "Himanshu Dogra"}
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Berry"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}
