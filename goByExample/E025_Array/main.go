/**
 * Example: Read an array
 */

package main

import (
	"fmt"
	"log"
)

func main() {
	var n int

	fmt.Println("Element number: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		log.Fatal(err)
	}

	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Println(i+1, " - th element: ")
		_, err := fmt.Scan(&array[i])

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(array)
}
