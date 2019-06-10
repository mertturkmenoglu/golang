/**
 * Error example
 */

package main

import "errors"
import "fmt"
import "math"

// Last return value is error by convention
// error is a built-in interface
func root(arg float64) (float64, error) {
	if arg < 0 {
		// If argument is negative, then return an arbitrary integer
		// And an error value. Basic error construction:
		return -1, errors.New("Number can not be negative")
	}

	// A nil value indicates no error happened
	return math.Sqrt(arg), nil
}

// You may create your own errors by implementing Error method
type negativeError struct {
	argument float64
	message  string
}

// Here is the implementation
func (e *negativeError) Error() string {
	return fmt.Sprintf("%f - %s", e.argument, e.message)
}

// Another function that may returns an error
func area(len float64) (float64, error) {
	if len <= 0 {
		// User defined error declaration syntax
		return -1, &negativeError{len, "Side length must be positive"}
	}
	return (len * len), nil
}

// Driver code
func main() {
	array := []int{4, 9, 16, 0, -1, 36}
	for _, number := range array {
		// Go type error check
		if returnValue, errorValue := root(float64(number)); errorValue != nil {
			fmt.Println("root:", errorValue)
		} else {
			fmt.Println("root:", returnValue)
		}
	}

	for _, i := range []int{-5, 10} {
		if r, e := area(float64(i)); e != nil {
			fmt.Println("area:", e)
		} else {
			fmt.Println("area:", r)
		}
	}

	_, e := area(42)
	if ae, ok := e.(*negativeError); ok {
		fmt.Println(ae.argument)
		fmt.Println(ae.message)
	}
}
