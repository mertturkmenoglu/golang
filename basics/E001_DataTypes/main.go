package main

import "fmt"

func main() {
	var integer int
	var float float64
	var boolean bool
	var str string

	fmt.Printf("%T %T %T %T\n", integer, float, boolean, str)
	fmt.Printf("%v %v %v %q\n", integer, float, boolean, str)
}
