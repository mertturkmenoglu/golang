package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(first, second string) bool {
	chars1 := strings.Split(first, "")
	chars2 := strings.Split(second, "")

	sort.Strings(chars1)
	sort.Strings(chars2)

	for i, e := range chars1 {
		if chars2[i] != e {
			return false
		}
	}

	return true
}

func main() {
	//noinspection SpellCheckingInspection
	result := isAnagram("emily", "mieyl")
	fmt.Println(result)

	fmt.Println(isAnagram("emily", "diana"))
}
