package main

import (
	"fmt"
	"strconv"
)

func main() {
	var integer = 42
	var floatingNumber = 3.14
	var sFirst = "1024"
	var sSecond = "61.456"

	fmt.Println(float64(integer))
	fmt.Println(int(floatingNumber))

	parsedInt, err := strconv.ParseInt(sFirst, 0, 64)

	if err != nil {
		fmt.Println("Parse error")
	} else {
		fmt.Println(parsedInt)
	}

	parsedFloat, err := strconv.ParseFloat(sSecond, 64)
	fmt.Println(parsedFloat)
}
