/**
 * Example 16: Variadic Functions
 *
 * Variadic function examples
 * Important points:
 *	- A variadic functions is a function that takes variable number
 * of arguments.
 */

package main

import "fmt"

/**
 * This function takes any number of arguments
 */
func sumOfIntegers(numbers ...int) {
	fmt.Print("You entered: ")
	fmt.Print(numbers, " ")

	sum := 0

	for _, tmp := range numbers {
		sum += tmp
	}

	fmt.Println("\nSum of the numbers: ", sum)
}

func main() {
	// Here calling with 1, 2 and 3 arguments
	sumOfIntegers(1)
	sumOfIntegers(1, 2)
	sumOfIntegers(1, 2, 3)

	// Calling function with a slice argument
	nums := []int{1, 2, 3, 4}
	sumOfIntegers(nums...)
}
