package main

import (
	"fmt"
	c "sectionThree/coffee"
	h "sectionThree/hammurabi"
	u "sectionThree/util"
)

func main() {
	options := map[int]string{
		1: "Coffee",
		2: "Hammurabi",
	}
	choice := u.MenuSelect(options)
	fmt.Print("\n\n\n")
	if choice == 1 {
		c.Coffee()
	} else if choice == 2 {
		h.Hammurabi()
	} else if choice != 0 {
		fmt.Printf("Invalid selection: %d\n", choice)
	}
}
