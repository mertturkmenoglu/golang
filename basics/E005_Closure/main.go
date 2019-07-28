package main

import "fmt"

func main() {
	n := 10

	increment := func() int {
		n += 1
		return n
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
