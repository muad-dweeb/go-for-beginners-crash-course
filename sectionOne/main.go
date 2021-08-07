package main

import (
	"bufio"
	"fmt"
	"os"
	"sectionOne/doctor"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}
}
