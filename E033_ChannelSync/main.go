/**
 * Channel synchronization example
 */
package main

import (
	"fmt"
	"time"
)

// Function will be invoked on a goroutine.
// Channel argument will be used to notify another
// goroutine.
func sleepForASec(records chan bool) {
	fmt.Println("Entered function")
	time.Sleep(time.Second)
	fmt.Println("Slept")

	// Send value to channel
	records <- true
}

// Driver code
func main() {
	// Make a channel
	records := make(chan bool, 1)

	//Invoke function on a goroutine
	go sleepForASec(records)

	// Block until receiving a notification from the channel
	<-records

	fmt.Println("After first part")

	// -----------------

	test := make(chan bool, 1)
	go sleepForASec(test)
	fmt.Println("It should not enter goroutine")

	fmt.Println("After all operations")
}
