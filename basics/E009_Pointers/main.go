package main

import "fmt"

// Just like in C
func main() {
	var a = 5
	var ptrA = &a

	var b = 3
	var ptrB = &b

	fmt.Println(a, b)

	swap(ptrA, ptrB)

	fmt.Println(a, b)
}

func swap(ptrA *int, ptrB *int) {
	var tmp = *ptrA
	*ptrA = *ptrB
	*ptrB = tmp
}
