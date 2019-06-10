/**
 * Channels example
 */

package main

import "fmt"

func main() {
	// Channels convey values between goroutines
	// You make a channel and declare its type, then send a
	// value from a goroutine get it from another goroutine

	// Create a channel
	myChannel := make(chan string)

	// Here we create an anonymous goroutine to send a string
	// to channel.
	go func() {
		myChannel <- "a string value goes to channel"
		myChannel <- "second string goes"
		myChannel <- "and the last one"
	}()

	// Take a string from channel and assign it to myStr
	myStr := <-myChannel

	// Print it
	fmt.Println(myStr)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)

	// If we write another one,
	// it will produce an error
	// fmt.Println(<-myChannel)

	// It seems that it works as a queue
}
