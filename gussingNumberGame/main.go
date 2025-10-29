package main

import (
	"fmt"
)

const maxBoundry = 100

var level int
var chances [3]int

// func runGame() {
// 	myGuess := binarySearch(0, 100)
//
// 	var counter int
//
// 	numOfChances := chances[level-1]
//
// 	for {
//
// 		fmt.Printf("my guess is: %v\n", myGuess)
//
// 		counter += 1
//
// 		fmt.Println("Is it correct? choose 0 if correct 1 if greater -1 if smaller")
// 		var input int
// 		fmt.Scan(&input)
//
// 		switch input {
// 		case 0:
// 			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.", counter)
// 			return
// 		case -1:
// 			myGuess = binarySearch(0, myGuess)
// 		case 1:
// 			myGuess = binarySearch(myGuess, 100)
// 		}
//
// 		if counter == numOfChances {
// 			fmt.Println("Game over you failed\n")
// 			return
// 		}
// 	}
// }

func main() {
	displayTut()

	initGame()

	runGame(0,100)
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

func runGame(left, right int)  {

	var counter int
	numOfChances := chances[level -1]

	for {
		if left > right {
			return
		}

		myGuess := ((left + right) / 2)
		fmt.Printf("my guess is: %v\n", myGuess)

		counter += 1

		fmt.Println("Is it correct? choose 0 if correct 1 if greater -1 if smaller")
		var input int
		fmt.Scan(&input)
		switch input {
		case 0:
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.", counter)
			return
		case -1:
			right = myGuess -1
		case 1:
			left = myGuess + 1
		}
		if counter == numOfChances {
			fmt.Println("Game over you failed\n")
			return
		}

	}

}
