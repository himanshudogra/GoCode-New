package main

import (
	"fmt"
)

func main() {

	var student [3]string
	fmt.Println(student)

	student[0] = "jack"
	student[1] = "ram"
	student[2] = "tom"
	fmt.Printf("student 1# %v\n", student[1])
}
