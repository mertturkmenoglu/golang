// Slice example
package main

import "fmt"

func main() {

	// Uninit slice
	slice := make([]int, 10)
	fmt.Println("Uninit: ", slice)

	fmt.Println("************")

	// Assign values
	for i := 0; i < 10; i++ {
		slice[i] = i
	}

	fmt.Println("Init: ", slice)

	fmt.Println("************")

	// Same as arrays, slice has len
	fmt.Println("Number of elements: ", len(slice))

	fmt.Println("************")

	// Append operation returns a new slice
	slice = append(slice, 10)
	fmt.Println("Appended: ", slice)
	fmt.Println("Number of elements: ", len(slice))

	// Make another slice
	other := make([]int, len(slice))

	// Copy(dest, src)
	copy(other, slice)
	fmt.Println("Other: ", other)

	fmt.Println("************")

	// Slice operator [low:high]
	// Includes low, excludes high
	newSlice := slice[2:5]
	fmt.Println("New: ", newSlice)

	// Slice operator is similar to Python version
	newSlice = slice[5:]
	fmt.Println("[5:] : ", newSlice)

	newSlice = slice[:3]
	fmt.Println("[:3] : ", newSlice)

	fmt.Println("************")

	// Create and init a slice
	initSlice := []string{"Winter", " is", " coming."}
	fmt.Println("House Stark: ", initSlice)

	fmt.Println("************")

	// Multi-dimensional slices
	// Length of the slices may vary
	twoDimSlice := make([][]int, 4)

	for i := 0; i < 4; i++ {
		sliceLen := i + 2
		twoDimSlice[i] = make([]int, sliceLen)
		for j := 0; j < sliceLen; j++ {
			twoDimSlice[i][j] = i + 2
		}
	}
	fmt.Println("Two dimensional slice: ", twoDimSlice)
}
