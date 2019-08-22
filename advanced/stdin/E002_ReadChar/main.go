package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter char: ")
	char, _, err := reader.ReadRune()

	if err != nil {
		log.Fatal(err)
	}

	// Print unicode value
	fmt.Println(char)

	// Print char
	fmt.Println(string(char))
}
