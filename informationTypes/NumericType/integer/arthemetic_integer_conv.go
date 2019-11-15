package main

import (
	"fmt"
)

func main() {

	var a int = 32
	var b int8 = 44

	fmt.Println(a + int(b)) // Even though int, int8 is of equivalent values but GO doesn't consider it as same, You need to explicitely convert the int8 to int type for arth operations.
}
