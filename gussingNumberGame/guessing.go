package main

import (
	"fmt"
	"math/rand"
)


var level int
var chances [3]int

func runGame() {
	targetNum := (rand.Intn(maxBoundry) + 1)

	var counter int

	numOfChances := chances[level-1]

	var userGuess int

	for {
		fmt.Printf("Enter your guess: ")
		fmt.Scan(&userGuess)
		counter += 1

		if userGuess == targetNum {
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.", counter)
			return
		}else if userGuess > targetNum {
			fmt.Printf("It's less than %v\n", userGuess)
		}else if userGuess < targetNum {
			fmt.Printf("It's greater than %v\n", userGuess)
		}

		if counter == numOfChances {
			fmt.Println("Game over you failed\n")
			return
		}
	}
}

func main() {
	displayTut()

	initGame()

	runGame()
}

func displayTut() {
	fmt.Println("Welcome to the number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println(" have 5 chances to guess the correct number.")
	fmt.Printf("Please select the difficulty level:\n 1. Easy( 10 chances )\n 2. Med ( 5 chances )\n 3. Hard (3 chances)\n")
	fmt.Print("Enter your choice: ")

	fmt.Scan(&level)

	fmt.Printf("\nGreat let's start our game\n")
}

func initGame() {
	chances[0] = 10
	chances[1] = 5
	chances[2] = 3
}
