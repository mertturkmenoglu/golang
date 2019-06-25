/**
 * Channel buffering example
 */

package main

import "fmt"

func main() {
	mychan := make(chan string)

	mychan <- "first"
	mychan <- "second"

	fmt.Println(<-mychan)
	fmt.Println(<-mychan)

}
