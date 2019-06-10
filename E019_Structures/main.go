/**
 * Example 19: Structures
 *
 * Basic structure example.
 */

package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

func main() {
	// Create and print a student
	fmt.Println(student{1, "Student1", "9/A"})

	// Create and print a student: Using field names
	fmt.Println(student{name: "Student2", id: 2, class: "10/A"})

	// Create and print a student: If you give no value for a field
	// it will be init to type's zero value
	fmt.Println(student{name: "Student3", class: "9/A"})

	// Create and print a student: Get a pointer to the struct
	fmt.Println(&student{id: 4, name: "Student4", class: "12/A"})

	// Accessing a field with a dot operator. Just like in C.
	student5 := student{5, "Student5", "11/A"}
	fmt.Println("Name: ", student5.name)

	// Accessing a field over a pointer to struct. You can use
	// dot operator to access.
	pStudent := &student5
	fmt.Println("ID: ", pStudent.id)

	// Rewrite field
	pStudent.id = 6
	fmt.Println("ID: ", pStudent.id)
}
