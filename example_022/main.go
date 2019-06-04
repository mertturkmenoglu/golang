/*
 * All zero n x n matrix example
 */

package main

import "fmt"

func main() {
	const n = 5
	var mtr [n][n]int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mtr[i][j] = 0
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(mtr[i][j], "\t")
		}
		fmt.Println()
	}
}
