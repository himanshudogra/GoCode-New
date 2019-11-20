package main

import (
	"fmt"
)

func main() {

	var identityMatrix [3][3]int // This can be re-written like

	identityMatrix[0] = [3]int{1, 0, 0} // var identityMatrix [3][3]int=[3][3] int {[3]int{1,0,0},[3]int{0,1,0},[3]int{0,0,1}}
	identityMatrix[1] = [3]int{0, 1, 0} // fmt.Println(identityMatrix)
	identityMatrix[2] = [3]int{0, 0, 1}

	fmt.Println(identityMatrix)
}
