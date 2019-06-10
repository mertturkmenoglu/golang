/**
 * Example 14: Functions
 *
 * Basic functions examples.
 * Important points:
 *	- Functions requires explicit returns
 *	- In function definition, if consecutive parameters
 * have same type, type name can be omitted and you can
 * write only the last paramaters type
 */

package main

import "fmt"

func helloWorld() {
	fmt.Println("Hello World")
}

func addTwoInt(first int, second int) int {
	return first + second
}

func addThreeInt(first, second, third int) int {
	return first + second + third
}

func main() {
	helloWorld()

	first := 1
	second := 2
	third := 3

	firstResult := addTwoInt(first, second)
	secondResult := addThreeInt(first, second, third)

	fmt.Println("First: ", firstResult)
	fmt.Println("Second: ", secondResult)
}
