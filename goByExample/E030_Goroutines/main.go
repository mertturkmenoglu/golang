/**
 * Goroutines example
 */

package main

import "fmt"

// We will try to run this function in a thread
// It takes a string and prints it for an arbitrary number of times
func printMessage(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println(message, ":", i)
	}
}

// Driver code
func main() {
	// A normal function call
	printMessage("Calling directly. Normal call")

	// Invoking in a thread
	go printMessage("Calling in a goroutine")

	// Anonymous function call in a goroutine
	go func(message string) {
		fmt.Println(message)
	}("go running")

	_, _ = fmt.Scanln()

	fmt.Println("Program ends")
}
