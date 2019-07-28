/**
 * Variable definitions example
 */

package main

import "fmt"

func main() {
	var num1 int16
	var str1 string
	var num2 float64

	var num3, num4 int8
	var str2, str3 string

	// var num5 int = 42
	// var str4 string = "answer"

	var num6 = 3.14

	num7 := 100
	num8 := 65535

	num9, str5 := 0, "zero"

	var (
		num10 = 1
		str6  = "string"
		num11 = 6.28
	)

	fmt.Println(num1)
	fmt.Println(str1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(str2)
	fmt.Println(str3)
	// fmt.Println(num5)
	// fmt.Println(str4)
	fmt.Println(num6)
	fmt.Println(num7)
	fmt.Println(num8)
	fmt.Println(num9)
	fmt.Println(str5)
	fmt.Println(num10)
	fmt.Println(str6)
	fmt.Println(num11)
}
