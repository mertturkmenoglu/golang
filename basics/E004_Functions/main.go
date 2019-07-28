package main

import "fmt"

type Person struct {
	name string
	id   int32
}

func (p Person) printInfo() {
	fmt.Println(p.name, p.id)
}

func addOne(number int) int {
	return number + 1
}

func addOneAndTwo(number int) (int, int) {
	return number + 1, number + 2
}

func main() {
	p := Person{
		name: "Emily",
		id:   1,
	}

	p.printInfo()

	result1 := addOne(5)
	fmt.Println(result1)

	result2, result3 := addOneAndTwo(5)
	fmt.Println(result2, result3)
}
