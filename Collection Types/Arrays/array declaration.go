package main

import (
	"fmt"
)

func main() {

	grades := [3]int{23, 34, 43} // can be declare like "grades:= [...]int{23,34,43}"
	fmt.Printf("%v,%T\n", grades, grades)
}
