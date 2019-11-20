package main

import (
	"fmt"
)

func main() {

	var student [5]string

	student[0] = "jack"
	student[1] = "ram"
	student[2] = "sapna"

	fmt.Printf("Numbers of students: %v\n", len(student))
}
