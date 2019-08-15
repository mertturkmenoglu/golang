package main

import "fmt"

func forEach(arr []int, f func(int)) {
	for _, e := range arr {
		f(e)
	}
}

func printInt(e int) {
	fmt.Println(e)
}

func main() {
	arr := []int{1, 2, 3}
	forEach(arr, printInt)
}
