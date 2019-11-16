package main

import (
	"fmt"
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeefinalcials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {

	var roles byte = isAdmin | canSeefinalcials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("IsAdmin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ?%v\n", isHeadquarters&roles == isHeadquarters)

}
