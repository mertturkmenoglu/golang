// Switch statement example
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// Classic switch statement
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		case 7:
			fmt.Println("Seven")
		case 8:
			fmt.Println("Eight")
		case 9:
			fmt.Println("Nine")
		default:
			fmt.Println("Unknown")
		}
	}

	// Enhanced version
	// You can separate cases with comma
	// They share a common operation
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 1, 2, 3, 4:
			fmt.Println("Less than five")
		case 5:
			fmt.Println("Equal to 5")
		case 6, 7, 8, 9:
			fmt.Println("Greater than five")
		default:
			fmt.Println("Unknown")
		}
	}

	// Enhanced version
	// Boolean expressions are valid for cases
	for i := 0; i < 10; i++ {
		switch {
		case i < 5:
			fmt.Println("Less than five")
		case i >= 5:
			fmt.Println("Greater than or equal to five")
		}
	}
}
