package main

import (
	"crashcourse/doctor"
	"fmt"
)

// func main() {
// 	// var saying string
// 	// saying = "Saluton mondo!!"
// 	saying := "Saluton mondo!!!"
// 	saySomething(saying)
// }

// func saySomething(whatToSay string) {
// 	fmt.Println(whatToSay)
// }

func main() {
	var whatToSay string
	whatToSay = doctor.Intro()
	fmt.Println(whatToSay)
}
