package main

import (
	"sectionTwo/scope"
)

var myVar = "This is a package level variable"

func main() {

	var blockVar = "This is the block level variable"

	scope.PrintMe(myVar, blockVar)
}
