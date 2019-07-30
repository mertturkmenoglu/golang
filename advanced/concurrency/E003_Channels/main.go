package main

import "fmt"

func oddSum(oddNumberChannel chan int, n int) {
	result := 0
	if n > 1 {
		for i := 1; i <= n; i++ {
			if i%2 == 1 {
				result += i
			}
		}
	}

	oddNumberChannel <- result
}

func evenSum(evenNumberChannel chan int, n int) {
	result := 0
	if n > 1 {
		for i := 1; i <= n; i++ {
			if i%2 == 0 {
				result += i
			}
		}
	}

	evenNumberChannel <- result
}

func main() {
	oddCh := make(chan int)
	evenCh := make(chan int)

	go oddSum(oddCh, 5)
	go evenSum(evenCh, 5)

	oddResult, evenResult := <-oddCh, <-evenCh

	fmt.Println(oddResult, " ", evenResult)
}
