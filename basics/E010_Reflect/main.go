package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int32
}

func main() {
	name := "Emily"
	fmt.Printf("%T\n", name)
	fmt.Println(reflect.TypeOf(name))

	age := 35
	fmt.Printf("%T\n", age)
	fmt.Println(reflect.TypeOf(age))

	p := Person{
		name: "emily",
		age:  35,
	}

	fmt.Printf("%T\n", p)
	fmt.Println(reflect.TypeOf(p))
}
