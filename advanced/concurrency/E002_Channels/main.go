package main

import "fmt"

func goRoutine(ch chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println("goRoutine function")
	}
	ch <- true
}

func main() {
	myChannel := make(chan bool)
	go goRoutine(myChannel)

	<-myChannel
	fmt.Println("End of the program")
}
