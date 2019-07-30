/**
 * Create identity matrix
 */

package main

import (
	"fmt"
	"log"
)

func main() {
	var n int

	fmt.Println("Matrix size: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		log.Fatal(err)
	}

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
