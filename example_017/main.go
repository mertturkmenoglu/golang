/**
 * Example 17: Recursive function example
 *
 * Quick Sort implementation
 */

package main

import "fmt"

func quickSort(array []int, begin int, end int) {
	if begin < end {
		partitionIndex := partition(array, begin, end)
		quickSort(array, begin, partitionIndex-1)
		quickSort(array, partitionIndex+1, end)
	}
}

func partition(array []int, begin int, end int) int {
	pivot := array[end]
	i := begin - 1
	for j := begin; j <= end-1; j++ {
		if array[j] <= pivot {
			i++
			tmp := array[i]
			array[i] = array[j]
			array[j] = tmp
		}
	}

	tmp := array[i+1]
	array[i+1] = array[end]
	array[end] = tmp
	return i + 1
}

func main() {
	array := []int{4, 2, 6, 4, 1, 3, 6, 7, 5, 8, 0, 9, 9, 10, 2}
	fmt.Println("Unsorted array: ", array)
	quickSort(array, 0, len(array)-1)
	fmt.Println("Sorted array: ", array)
}
