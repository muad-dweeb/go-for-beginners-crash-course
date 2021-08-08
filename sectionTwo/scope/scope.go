package scope

import "fmt"

var PackageVar = "This is a package level variable in PackageOne"

func PrintMe(s1, s2 string) {
	fmt.Println(s1, 2, PackageVar)
}
