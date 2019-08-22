package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter something: ")
		text, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		text = strings.Replace(text, "\r\n", "", -1)
		fmt.Println(text)
	}
}
