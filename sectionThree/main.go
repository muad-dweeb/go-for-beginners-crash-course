package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/eiannone/keyboard"
)

func menu_select(input_map map[int]string) int {
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

func coffee() {
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

	menu_select(menu)
}

func hammurabi() {
	fmt.Println("Not implemented")
}

func main() {
	options := map[int]string{
		1: "Coffee",
		2: "Hammurabi",
	}
	choice := menu_select(options)
	fmt.Print("\n\n\n")
	if choice == 1 {
		coffee()
	} else if choice == 2 {
		hammurabi()
	} else if choice != 0 {
		fmt.Printf("Invalid selection: %d\n", choice)
	}
}
