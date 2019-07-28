package main

import (
	"errors"
	"fmt"
	"math"
)

func getSqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, errors.New("negative number passed to getSqrt()")
	}

	return math.Sqrt(number), nil
}

func main() {
	result, err := getSqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = getSqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
