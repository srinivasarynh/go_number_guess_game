package main

import (
	"fmt"
	"math/rand"
)


func main() {

	fmt.Println("welcome to the guess the number..")
	fmt.Println("I'm thinking of a number between 1 and 100")

	playAgain := true
	for playAgain {
		playGame()
		playAgain = askPlayAgain()
	}

	fmt.Println("thanks for playing...")
}

func playGame() {
	secretNumber := rand.Intn(100) + 1 
	attempts := 0
	maxAttempts := 10 
	fmt.Printf("\nyou have %d attempts to guess the number!\n", maxAttempts)
	for attempts < maxAttempts {
		attempts++
		fmt.Printf("\nattempt %d/%d - enter your guess: ", attempts, maxAttempts)

		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("please enter valid number")
			attempts--
			continue
		}
		if guess < 1 || guess > 100 {
			fmt.Println("please guess a number bewteen 1 and 100")
			attempts--
			continue
		}

		if guess == secretNumber {
			fmt.Printf("congratulations you guessed it in %d attempts\n", attempts)
			return
		} else if guess < secretNumber {
			fmt.Println("too low try a higher number.")
		} else {
			fmt.Println("too high try a lower number")
		}

		remainingAttempts := maxAttempts - attempts
		if remainingAttempts > 0 {
			fmt.Printf("you have %d attempts remaining. \n", remainingAttempts)
		}
	}
	fmt.Printf("game over! the number was %d better luck next time!\n", secretNumber)
}


func askPlayAgain() bool {
	fmt.Print("\nwould you like to play again? (y/n): ")
	var response string
	fmt.Scanf("%s", &response)

	return response == "y" || response == "Y" || response == "yes" || response == "Yes"
}
