package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Making sure that we have the file all the time
	file, err := os.Create("E002_FileIO_text.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	_, _ = file.WriteString("First line\n")
	_, _ = file.WriteString("Second line")

	// Now read file
	stream, err := ioutil.ReadFile("E002_FileIO_text.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println(readString)
}
