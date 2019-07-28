package main

import "fmt"

type Person struct {
	fName   string
	lName   string
	address Address
}

type Address struct {
	street string
	no     int32
}

func (p Person) printDetails() {
	fmt.Println(p)
}

func (a Address) printDetails() {
	fmt.Println(a)
}

func main() {
	add := Address{
		street: "GoodStreet",
		no:     10,
	}

	person := Person{
		fName:   "Emily",
		lName:   "Smith",
		address: add,
	}

	add.printDetails()
	person.printDetails()
}
