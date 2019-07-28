package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var message = "Hello World"

	fmt.Println(message)
	fmt.Println(reflect.TypeOf(message))

	fmt.Println(len(message))

	for char := 65; char < 91; char++ {
		fmt.Print(string(char), ":", char)
	}

	name := "emily"
	fmt.Println(strings.ToUpper(name))

	surname := "SMITH"
	fmt.Println(strings.ToLower(surname))

	fmt.Println(strings.HasPrefix(name, "em"))
	fmt.Println(strings.HasSuffix(name, "ly"))

	var characters = []string{"d", "o", "c", "t", "o", "r"}
	fmt.Println(strings.Join(characters, ""))

	fmt.Println(strings.Repeat("New ", 4), "York")
	fmt.Println(strings.Contains("Emily", "il"))
	fmt.Println(strings.Index("Emily", "il"))
	fmt.Println(strings.Count("Emily is the best", "e"))
	fmt.Println(strings.Replace("Hello there. General Kenobi!", "e", "E", 100))

	str := "Sing us a song you are the piano man"
	var charArray = strings.Split(str, " ")

	for index, value := range charArray {
		fmt.Println("Index : ", index, "value : ", value)
	}
}
