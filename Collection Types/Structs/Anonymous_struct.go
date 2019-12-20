package main

import (
	"fmt"
)

func main() {

	aDoctor := struct{ name string }{name: "Lisa Roy"}
	fmt.Println(aDoctor)
}
