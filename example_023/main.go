/*
 * Sum of two n x n matrices
 */

package main

import "fmt"

func main() {
	const n = 5

	var mtrA [n][n]int
	var mtrB [n][n]int
	var mtrC [n][n]int

	factor := 1

	// Assign values
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mtrA[i][j] = i*n + j + 1
			mtrB[i][j] = factor
			if factor < 128 {
				factor *= 2
			} else {
				factor = 1
			}
		}
	}

	// Find the sum
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mtrC[i][j] = mtrA[i][j] + mtrB[i][j]
		}
	}

	// Print
	fmt.Print(mtrA)
	fmt.Println(mtrB)
	fmt.Println(mtrC)
}
