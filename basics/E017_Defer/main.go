package main

import "fmt"

// noinspection GoDeferInLoop
func printNumbers(numbers ...int) {
	defer fmt.Println("Ended")
	for _, e := range numbers {
		defer fmt.Println(e)
	}
	defer fmt.Println("Started")
}

func main() {
	printNumbers(1, 2, 3, 4, 5)
}
