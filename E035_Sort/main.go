/**
 * Sorting example
 */

package main

import "fmt"
import "sort"

func main() {
	students := []string{"john", "amelia", "david"}
	sort.Strings(students)
	fmt.Println("String sort: ", students)

	ids := []int{443, 123, 591}
	sort.Ints(ids)
	fmt.Println("Integer sort: ", ids)

	isSorted := sort.IntsAreSorted(ids)
	fmt.Println("Sorted: ", isSorted)
}
