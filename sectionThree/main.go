package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	menu := map[int]string{
		1: "Macchiato",
		2: "Dirty Chai",
		3: "Mocha",
		4: "Espresso",
		5: "Cortado",
		6: "Flat White",
	}

	fmt.Println("MENU")
	fmt.Println("----")

	// Extract and sort the map keys
	keySlice := make([]int, 0, len(menu))
	for num := range menu {
		keySlice = append(keySlice, num)
	}
	sort.Ints(keySlice)

	// Print the menu
	for _, item := range keySlice {
		fmt.Println(item, "-", menu[item])
	}
	fmt.Println("Q - Quit the program")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal((err))
		}

		if key == keyboard.KeyEsc || char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", menu[i]))
	}

	fmt.Println("Program exiting...")

	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Print("-> ")
	// 	userInput, _ := reader.ReadString('\n')
	// 	userInput = strings.TrimSpace(userInput)
	// 	if userInput == "quit" {
	// 		break
	// 	} else {
	// 		fmt.Println(userInput)
	// 	}
	// }
}
