package main

import "fmt"

func forEach(arr []int, n int, f func(e int)) {
	for i := 0; i < n; i++ {
		f(arr[i])
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := 5

	forEach(arr, n, func(e int) {
		fmt.Print(e, "\t")
	})

	fmt.Println()

	forEach(arr, n, func(e int) {
		fmt.Print(e*e, "\t")
	})
}
