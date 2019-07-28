/**
 * Channel buffering example
 */

package main

import "fmt"

func main() {
	myChannel := make(chan string)

	myChannel <- "first"
	myChannel <- "second"

	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)

}
