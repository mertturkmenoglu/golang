/**
 * Example: Read an array
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Println("Element number: ")
	fmt.Scan(&n)

	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Println(i+1, " - th element: ")
		fmt.Scan(&array[i])
	}

	fmt.Println(array)
}
