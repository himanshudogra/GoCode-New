package main

import (
	"fmt"
)

const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // 1 times 2, 10 times (10*iota = 10*1 = 10), total value = (1*2)^10 = 1024
	MB                    // 1 times 2, 20 times (10*iota = 10*2 = 20) , total value = (1*2)^20 = 1048576
	GB                    // 1 times 2, 30 times (10*iota = 10*3 = 30), total value = (1*2)^30 = 1073741824
	TB
	PB
	EB
	ZB
	YB
)

func main() {

	filesize := 4000000000.
	fmt.Printf("%.2fGB\n", filesize/GB) // 4000000000/1073741824 = 3.73 (maximum upto 2 decimal with GB unit)
}
