// Constants
package main

import "fmt"
import "math"

// Constant string (String literal)
const str = "Doctor Who?"

func main() {
	fmt.Println(str)

	const n = 1000
	const num = 3e10 / n

	// As scientific notation
	fmt.Println(num)
	// As 64 bit integer
	fmt.Println(int64(num))

	fmt.Println(math.Sin(n))
}
