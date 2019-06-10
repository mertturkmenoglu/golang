// Switch example - 3
package main

import "fmt"
import "time"

func main() {
	currTime := time.Now().Hour()

	switch {
	case currTime < 12:
		fmt.Println("Before noon")
	case currTime == 12:
		fmt.Println("It is noon")
	case currTime > 12:
		fmt.Println("After noon")
	}
}
