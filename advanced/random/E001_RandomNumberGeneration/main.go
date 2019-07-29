package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	// Every time, numbers will be the same
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}

	// Every time, numbers will be the same
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Float64())
	}

	// Now numbers will be pseudo-random
	rand.Seed(time.Now().Unix())
	number := randomRange(1, 20)

	fmt.Println(number)
}
