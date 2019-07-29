package main

import (
	"fmt"
	"sort"
)

type SortByLength []string

func (s SortByLength) Len() int {
	return len(s)
}

func (s SortByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	intNumbers := []int{11, 12, 9, 10}
	sort.Ints(intNumbers)
	fmt.Println("Integer numbers:", intNumbers)

	floatNumbers := []float64{9.71, 42.67, 2.7, 3.14}
	sort.Float64s(floatNumbers)
	fmt.Println("Float numbers:", floatNumbers)

	strings := []string{"Amelia", "Clara", "Billie"}
	sort.Strings(strings)
	fmt.Println("Strings:", strings)

	result := sort.Float64sAreSorted(floatNumbers)
	fmt.Println("Sorted:", result)

	// Special sort
	names := []string{"Emily", "Diana", "Barbara", "Selina"}
	sort.Sort(SortByLength(names))
	fmt.Println(names)
}
