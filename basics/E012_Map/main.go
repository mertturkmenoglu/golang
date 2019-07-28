package main

import "fmt"

type Coordinate struct {
	x, y float64
}

func main() {
	people := map[string]int{"Emily": 35, "Diana": 37, "Barbara": 22}
	fmt.Println(people)
	fmt.Println(people["Emily"])

	students := make(map[string]int)
	fmt.Println(students)
	students["Jennifer"] = 20
	students["Kyle"] = 18
	students["Bruce"] = 21

	fmt.Println(students)
	students["Kyle"] = 19

	fmt.Println(students)
	fmt.Println("Student age:", students["Jennifer"])

	cities := make(map[string]string)
	cities["Rome"] = "Italy"
	cities["London"] = "England"
	cities["Paris"] = "France"

	fmt.Println(cities)
	delete(cities, "Paris")
	fmt.Println(cities)

	movies := make(map[string]int)
	movies["LOTR"] = 3
	movies["Terminator"] = 5
	movies["Titanic"] = 1

	fmt.Println(movies)

	val1, status1 := movies["Terminator"]
	fmt.Println("Value:", val1, "Status:", status1)

	val2, status2 := movies["The Shining"]
	fmt.Println("Value:", val2, "Status:", status2)

	locations := map[string]Coordinate{
		"Park":   {56.77954, -96.15742},
		"Office": {44.47102, -111.42108},
	}

	fmt.Println(locations)
}
