// If else example
package main

import "fmt"

func main() {
	var n int
	n = 50

	// If without else branch
	for i := 0; i < n; i++ {
		if i%2 == 0 && i%3 == 0 {
			fmt.Println(i, " can be divided by 6")
		}
	}

	// If with a else branch
	if n%2 == 0 {
		fmt.Println("n is even")
	} else {
		fmt.Println("n is odd")
	}

	n = -1
	if n < 0 {
		fmt.Println("n is negative")
	} else if n == 0 {
		fmt.Println("n is zero")
	} else {
		fmt.Println("n is positive")
	}
}
