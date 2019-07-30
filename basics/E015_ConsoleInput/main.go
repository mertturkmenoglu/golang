package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var number float64

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your sentence: ")
	word, err := reader.ReadString('\n')
	word = strings.Replace(word, "\n", "", -1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter a number: ")
	_, _ = fmt.Scan(&number)

	fmt.Println(word, " - ", number)
}
