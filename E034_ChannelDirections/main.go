/**
 * Channel directions example
 */

package main

import "fmt"

func sendMessage(msgChan chan<- string, msg string) {
	msgChan <- msg
}

func recieveMessage(msgChan <-chan string) {
	fmt.Println(<-msgChan)
}

func main() {
	msgChan := make(chan string, 1)
	sendMessage(msgChan, "Message1")
	recieveMessage(msgChan)

	sendMessage(msgChan, "Message2")
	recieveMessage(msgChan)
}
