package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{SpeedKPH: 45.09, CanFly: true}
	a := Animal{Name: "Tiger", Origin: "India"}
	fmt.Println(a.Name)
	fmt.Println(b.CanFly)

}
