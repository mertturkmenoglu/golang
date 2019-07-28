/**
 * Example 18: Pointers
 *
 * Pointer example: A basic swap function
 * Important points:
 *	- Pointers operates like in C
 *	- We use '*' as deaddress operator and '&' as address operator
 */

package main

import "fmt"

func swap(first, second *int) {
	tmp := *first
	*first = *second
	*second = tmp
}

func swapDoesNotWork(first, second int) {
	tmp := first
	first = second
	second = tmp
}

func main() {
	first := 3
	second := 5

	fmt.Println("Before: ", first, second)

	swapDoesNotWork(first, second)
	fmt.Println("After unsuccessful swap: ", first, second)

	swap(&first, &second)
	fmt.Println("After successful swap: ", first, second)
}
