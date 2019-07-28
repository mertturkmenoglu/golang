/**
 * Multiple return values example
 */

package main

import "fmt"

func evenOdd(number int) (string, int) {
	if number%2 == 0 {
		return "even", 0
	}
	return "odd", 1
}

func main() {
	twenty := 20
	resultStr, resultInt := evenOdd(twenty)
	fmt.Println("String:", resultStr)
	fmt.Println("Integer:", resultInt)
}
