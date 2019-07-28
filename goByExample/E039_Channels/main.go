/**
 * Channel example
 */

package main

import "fmt"

func createMessage(numbers chan int, times int) {
	for i := 0; i < times; i++ {
		numbers <- i
	}
}

func printMessage(numbers chan int, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(<-numbers)
	}
}

func main() {
	numbers := make(chan int)
	times := 1000
	go createMessage(numbers, times)
	go printMessage(numbers, times)

	fmt.Scanln()
}
