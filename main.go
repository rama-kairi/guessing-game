package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a random number between 1-100
	rand.Seed(time.Now().Unix())
	secret := rand.Intn(100)

	// Print the secret number
	fmt.Println(secret)

	// Store a attempts counter
	attempt := 10

	for attempt > 0 {

		// TODO: Read the input code in a Function
		// Read the input from stdin (terminal)
		fmt.Print("Enter a number: ")
		var input int

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error reading input")
			return
		}

		attempt = attempt - 1

		// TODO: Replace the following if-else with a switch statement
		// if the Input is less than the secret number -> print "Input is too low"
		if input < secret {
			fmt.Printf("Input is too Low (attempts left: %d), try again \n", attempt)
		} else if input > secret {
			fmt.Printf("Input is too High (attempts left: %d), try again \n", attempt)
		} else {
			fmt.Printf("You guessed the number %d in %d attempts \n", secret, attempt)
			return
		}

		// Reduce the attempts counter
		attempt--
	}

	// Need to handle some edge cases
	- // TODO: Users should not be able to enter a number greater than 100 and less than 1
	- // TODO: If user enter any non integer value, the program should not crash and should ask for input again with warning message
}
