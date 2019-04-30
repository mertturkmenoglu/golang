package main

import "fmt"
import "math"

func main() {
	// Init loop index before loop
	i := 1
	for i <= 10 {
		fmt.Println(math.Pow(float64(i), 2))
		i++
	}

	// Classical for loop
	for j := 0; j <= 10; j++ {
		fmt.Println(math.Pow(float64(i), 3))
	}

	// While(1)
	for {
		fmt.Println("While(1)")
		break
	}

	// Golang has "continue" but it is not recommended to use
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
