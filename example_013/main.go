package main

import "fmt"

func main() {
	numbers := []int{12, 75, 44, 1, 27, 59, 42}
	sum := 0

	// The index and value.
	for _, num := range numbers {
		sum += num
	}

	fmt.Println("Sum of the numbers:", sum)

	for i, num := range numbers {
		if num%2 == 0 {
			fmt.Println("A[", i, "]: ", numbers[i])
		}
	}

	// It can also iterate over map key-values, map keys and strings
}
