// Array example
package main

import "fmt"

func main() {
	// Uninitialized array
	var array [10]int
	fmt.Println("Uninitialized array: ", array)

	// Init now
	for i := 0; i < 10; i++ {
		array[i] = i
	}

	// Print array, print element
	fmt.Println("Array: ", array)
	fmt.Println("Array[1]: ", array[1])

	// Print array size
	fmt.Println("Number of elements: ", len(array))

	// Init array
	array2 := [3]string{"england", "turkey", "russia"}
	fmt.Println(array2)

	const n = 2

	// !Important: n must be constant
	var mtr [n][n]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				mtr[i][j] = 1
			} else {
				mtr[i][j] = 0
			}
		}
	}
	fmt.Println(mtr)
}
