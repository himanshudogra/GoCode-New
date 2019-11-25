package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice to all the elements
	c := a[3:]  // slice from 3th index(which is the 4th element) until end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice from 4th, 5th, and 6th elements

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
