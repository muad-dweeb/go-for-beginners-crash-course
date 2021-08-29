package hammurabi

import (
	"fmt"
	"sectionThree/hammurabi/game"
)

func Hammurabi() {
	playAgain := true

	for playAgain {
		game.Play()
		playAgain = game.GetYesOrNo("Would you like to play again (y/n)?")
	}

	fmt.Println("")
	fmt.Println("Goodbye.")
}
