/*
 * Go examples: Find minimum integer in an array
 */

package main

import (
	"fmt"
	"log"
)

func main() {
	var array [10]int

	for i := 0; i < 10; i++ {
		fmt.Println("Element: ")
		_, err := fmt.Scan(&array[i])

		if err != nil {
			log.Fatal(err)
		}
	}

	var minIndex int
	minIndex = 0

	for i := 1; i < 10; i++ {
		if array[i] < array[minIndex] {
			minIndex = i
		}
	}

	fmt.Println(array[minIndex])
}
