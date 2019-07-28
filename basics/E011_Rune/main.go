package main

import (
	"fmt"
	"reflect"
)

func main() {
	runeVar := 'E'

	fmt.Println(runeVar)
	fmt.Printf("%d\n", runeVar)
	fmt.Printf("%c\n", runeVar)
	fmt.Println(reflect.TypeOf(runeVar))
}
