package main

import (
	"fmt"
	"math/rand"
	"time"
)

func askStdnInput(min, max int) (int, bool) {
	// Read the input from stdin (terminal)
	fmt.Print("Enter a number: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Input is not a number, try again")
		return input, false
	}

	if input < min || input > max {
		fmt.Println("Input is out of range, it should be between 1-100")
		return input, false
	}
	return input, true
}

func getOnlyIntFromStdin(min, max int) int {
	// Read the input from stdin (terminal)
	fmt.Print("Enter a number: ")

	var input int
	_, err := fmt.Scan(&input)

	if input < min || input > max {
		fmt.Println("Input is out of range, it should be between 1-100")
		return getOnlyIntFromStdin(min, max)
	}

	if err != nil {
		fmt.Println("Input is not a number, try again")
		return getOnlyIntFromStdin(min, max)
	}

	return input
}

func main() {
	// Generate a random number between 1-100
	rand.Seed(time.Now().Unix())
	secret := rand.Intn(100)

	// Print the secret number
	fmt.Println(secret)

	// Store a attempts counter
	attempt := 9

	for attempt >= 0 {

		input := getOnlyIntFromStdin(1, 100)

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
		if attempt == 0 {
			fmt.Println("You have exhausted all your attempts, the secret number was ", secret)
			return
		}

		// Use switch case

		switch {
		case input < secret:
			{
				fmt.Printf("Input is too Low (attempts left: %d), try again \n", attempt)
			}
		case input > secret:
			{
				fmt.Printf("Input is too High (attempts left: %d), try again \n", attempt)
			}
		default:
			{
				fmt.Printf("You guessed the number %d in %d attempts \n", secret, attempt)
				return
			}
		}

		// Reduce the attempts counter

		attempt--

	}

	// Need to handle some edge cases
	// TODO: If user enter any non integer value, the program should not crash and should ask for input again wth
}
