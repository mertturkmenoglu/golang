package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func executeCommand(os string) {
	var command string

	if os == "linux" {
		command = "ls"
	} else {
		log.Fatal("Windows is not supported.")
	}

	out, err := exec.Command(command).Output()

	if err != nil {
		log.Fatal(err)
	}

	output := string(out[:])
	fmt.Println(output)
}

func main() {
	executeCommand(runtime.GOOS)
}
