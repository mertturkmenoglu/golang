package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("text.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, _ = file.WriteString("Go file input / output")
}
