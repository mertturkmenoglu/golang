package main

import "fmt"

/**
 * Variables example
 */
//noinspection ALL
func main() {
	// Declaring and initializing a variable
	var str1 = "Hello"
	fmt.Println(str1)

	// Declaring multiple variables
	var int1, int2 int = 1, 2
	fmt.Println(int1, int2)

	// Declaring Boolean
	var isTrue = true
	fmt.Println(isTrue)

	// Not initializing. Default is 0 for integer
	var int3 int
	fmt.Println(int3)

	// Declaring and initializing: Shorthand syntax
	str2 := "World"
	fmt.Println(str2)
}
