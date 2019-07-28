package main

import "fmt"

func main() {
	// Creating an empty map
	// Map holds a value as key-value pair
	words := make(map[string]string)

	words["stark"] = "Winter is coming"
	words["lannister"] = "Hear me roar"
	fmt.Println("Houses and words: ", words)

	fmt.Println("************")

	// Number of elements
	fmt.Println("House number: ", len(words))

	fmt.Println("************")

	// Get a value
	starkMotto := words["stark"]
	fmt.Println("House Stark Motto: ", starkMotto)

	fmt.Println("************")

	// Add another house
	words["mormont"] = "Here we stand"
	fmt.Println("Houses and words: ", words)

	// Delete House Mormont
	delete(words, "mormont")
	fmt.Println("Houses and words: ", words)

	fmt.Println("************")

	// Create and initialize a map

	ages := map[string]int{"Alex": 23, "David": 18, "Lyanna": 22}
	fmt.Println("Person - Age: ", ages)
}
