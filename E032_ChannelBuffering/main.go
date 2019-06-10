package main

import "fmt"

func main() {
	// Make a channel with buffering up to 2 values
	stringChannel := make(chan string, 2)

	// Channel is buffered so we can send values without receiving
	stringChannel <- "first string"
	stringChannel <- "second string"

	// Receive it later
	fmt.Println(<-stringChannel)
	fmt.Println(<-stringChannel)
}
