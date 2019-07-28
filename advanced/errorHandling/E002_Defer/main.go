package main

import "fmt"

func printMessage(s string) {
	fmt.Println(s)
}
func printMessage2(s string) {
	fmt.Println(s)
}

func main() {
	defer printMessage("First message")
	printMessage2("Second message")
}
