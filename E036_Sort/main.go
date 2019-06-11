/**
 * Custom sort example
 */

package main

import "sort"
import "fmt"

// We will call sort.Sort generic function. But we need to implement
// sort.Interface.Len
// sort.Interface.Less
// sort.Interface.Swap
// So we define our type then implement these methods on this type.
// sortByLength is our new type. It is actually nothing but a string array
type sortByLength []string

// Implement Len() method
// It returns element count of the array
// In this example, we do not need to change it's functionality
func (s sortByLength) Len() int {
	return len(s)
}

// Implement Swap() method
// It swaps two elements according to given indices
// In this example, we do not need to change it's functionality
func (s sortByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Implement Less() method
// It is used to determine which element is the small one
// By default, strings are sorted in order to their
// lexicographical order but we want to sort them by their lengths
// So we need a function to compare two individual elements
// Less() method compares two given element
func (s sortByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// Driver code
func main() {
	characters := []string{"john wick", "neo", "agent smith", "baba yaga"}
	fmt.Println(characters)

	sort.Sort(sortByLength(characters))

	fmt.Println(characters)
}
