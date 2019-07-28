// Switch example - 2
package main

import "fmt"

func main() {
	printType := func(num interface{}) {
		switch typeOfVar := num.(type) {
		case bool:
			fmt.Println("Variable is a boolean")
		case int:
			fmt.Println("Variable is an integer")
		case string:
			fmt.Println("Variable is a string")
		default:
			// C type formatted output
			fmt.Printf("Unknown type: %T\n", typeOfVar)
		}
	}

	printType("I must be a string")
	printType(false)
	printType(42)
	printType(3.14)
}
