package main

import "fmt"

/*
     The function is intended to calculate the factorial of a number,
     but it is currently calculating the sum of numbers from 1 to n instead of the factorial.

    // This is the correct implementation of the factorial function
    func factorial(n int) int {
		result := 1
		for i := 1; i <= n; i++ {
			result *= i // Multiplying the result by the current number i
		}
		return result
	}

*/

func factorial(n int) int {
	result := 1.00
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

// Function to print the factorial of a number
func printFactorial() {
	num := 5
	fmt.Printf("The factorial of %d is: %d\n", num, factorial(num))
}

// The wrong function name is used here.
// It should be printFactorial instead of printfactorial
func main() {
	printfactorial()
}
