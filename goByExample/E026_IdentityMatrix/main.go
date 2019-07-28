/**
 * Create identity matrix
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Println("Matrix size: ")
	fmt.Scan(&n)

	mtr := make([][]int, n)

	for i := 0; i < n; i++ {
		mtr[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				mtr[i][j] = 1
			} else {
				mtr[i][j] = 0
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(mtr[i])
	}
}
