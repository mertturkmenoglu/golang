package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	words := strings.Fields(str)

	for i, w := range words {
		words[i] = func(str string) string {
			chars := []rune(str)

			for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
				chars[i], chars[j] = chars[j], chars[i]
			}

			return string(chars)
		}(w)
	}

	return strings.Join(words, " ")
}
func main() {
	reversed := reverseWords("emily is the best")
	fmt.Println(reversed)
}
