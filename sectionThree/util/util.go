package util

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/eiannone/keyboard"
)

func MenuSelect(input_map map[int]string) int {
	var i int
	fmt.Println("MENU")
	fmt.Println("----")

	// Extract and sort the map keys
	keySlice := make([]int, 0, len(input_map))
	for num := range input_map {
		keySlice = append(keySlice, num)
	}
	sort.Ints(keySlice)

	// Print the menu
	for _, item := range keySlice {
		fmt.Println(item, "-", input_map[item])
	}
	fmt.Println("Q - Quit the program")

	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		log.Fatal((err))
	}

	if key == keyboard.KeyEsc || char == 'q' || char == 'Q' {
		i = 0
		fmt.Println("Program exiting...")
		return i
	} else {
		i, _ := strconv.Atoi(string(char))
		fmt.Printf("\nYou chose %s\n", input_map[i])
		return i
	}
}
