/**
 * Channel directions example
 */

package main

import "fmt"

func sendMessage(msgChan chan<- string, msg string) {
	msgChan <- msg
}

func receiveMessage(msgChan <-chan string) {
	fmt.Println(<-msgChan)
}

func main() {
	msgChan := make(chan string, 1)
	sendMessage(msgChan, "Message1")
	receiveMessage(msgChan)

	sendMessage(msgChan, "Message2")
	receiveMessage(msgChan)
}
