package main

import "fmt"

func main() {
	numbers := []int{2, 3, 4}
	sum := 0

	// Ignore index, just get value
	for _, value := range numbers {
		sum += value
	}

	fmt.Println("sum:", sum)

	for i, value := range numbers {
		fmt.Println("Index:", i, "Value:", value)
	}

	characters := map[int]string{0: "emily", 1: "diana", 2: "barbara"}

	for key, value := range characters {
		fmt.Printf("%d -> %s\n", key, value)
	}
	for key := range characters {
		fmt.Println("Character:", key)
	}

	for index, character := range "Mind the gap" {
		fmt.Println(index, character)
	}
}
