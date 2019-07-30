/**
 * Goroutines example
 */

package main

import "fmt"

func printMessage(message string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(message, "-", i)
	}
}

func main() {
	go printMessage("Goroutine", 100)
	go printMessage("Emily", 100)
	go printMessage("Diana", 100)
	go printMessage("Babushka", 100)

	_, _ = fmt.Scanln()
}
