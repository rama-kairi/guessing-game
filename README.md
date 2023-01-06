# **Number Guessing Game**

This is a simple number guessing game where the user tries to guess a randomly generated number between 1 and 100. The user has 10 attempts to guess the correct number.

## **Dependencies**

- "fmt"
- "math/rand"
- "time"

## **Usage**

1. Run the **`main`** function.
2. Enter a number between 1 and 100.
3. The app will tell you if your guess is too high or too low.
4. If you guess the correct number within 10 attempts, the app will tell you how many attempts it took.
5. If you are unable to guess the correct number within 10 attempts, the app will tell you the correct number.

## **Customization**

You can customize the range of the secret number by modifying the **`min`** and **`max`** variables in the **`getOnlyIntFromStdin`** function. You can also adjust the number of attempts by modifying the **`attempt`** variable in the **`main`** function.
