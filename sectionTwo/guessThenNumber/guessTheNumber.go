package guessTheNumber

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready"

func instruct(firstNumber int, secondNumber int, subtraction int, answer int) {
	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")
	fmt.Println("Type a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')
	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is", answer)
}

func Run() {
	// seed the RNG
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	instruct(firstNumber, secondNumber, subtraction, answer)
}