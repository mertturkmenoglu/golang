/**
 * Example 20: Methods
 *
 * Methods examples.
 */
package main

import "fmt"
import "math"

// Struct definition
type circle struct {
	radius float32
}

// Receiver type: circle*
func (c *circle) area() float32 {
	return c.radius * c.radius * math.Pi
}

// Receiver type: value
func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}

// Driver code to test methods
func main() {
	c := circle{radius: 1}

	fmt.Println("Circle area: ", c.area())
	fmt.Println("Circle perimeter:", c.perimeter())

	pointerC := &c
	fmt.Println("Circle area: ", pointerC.area())
	fmt.Println("Circle perimeter:", pointerC.perimeter())
}
