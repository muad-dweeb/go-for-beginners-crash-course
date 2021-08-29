package coffee

import (
	"log"
	u "sectionThree/util"

	"github.com/eiannone/keyboard"
)

func Coffee() {
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
	u.MenuSelect(menu)
}
