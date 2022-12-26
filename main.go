package main

import (
	"fmt"
	"math/rand"
	"time"
)

func askStdnInput(min, max int) int {
	// Read the input from stdin (terminal)
	fmt.Print("Enter a number: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input")
		return 0
	}

	if input < min || input > max {
		fmt.Println("Input is out of range, it should be between 1-100")
		return 0
	}
	return input
}

func main() {
	// Generate a random number between 1-100
	rand.Seed(time.Now().Unix())
	secret := rand.Intn(100)

	// TODO: Handle the Attempt issue for out of range input cases

	// Print the secret number
	fmt.Println(secret)

	// Store a attempts counter
	attempt := 10

	for attempt > 0 {

		input := askStdnInput(1, 100)

		// // if the Input is less than the secret number -> print "Input is too low"

		// Using if else

		// if input < secret {
		// 	fmt.Printf("Input is too Low (attempts left: %d), try again \n", attempt)
		// } else if input > secret {
		// 	fmt.Printf("Input is too High (attempts left: %d), try again \n", attempt)
		// } else {
		// 	fmt.Printf("You guessed the number %d in %d attempts \n", secret, attempt)
		// 	return
		// }

		// Use switch case
		switch {
		case input < secret:
			{
				fmt.Printf("Input is too Low (attempts left: %d), try again \n", attempt-1)
			}
		case input > secret:
			{
				fmt.Printf("Input is too High (attempts left: %d), try again \n", attempt-1)
			}
		default:
			{
				fmt.Printf("You guessed the number %d in %d attempts \n", secret, attempt-1)
				return
			}
		}

		// Reduce the attempts counter
		attempt--
	}

	// Need to handle some edge cases
	// TODO: If user enter any non integer value, the program should not crash and should ask for input again wth
}
