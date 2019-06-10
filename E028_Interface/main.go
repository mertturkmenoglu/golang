/**
 * Interface example
 */

package main

import "fmt"

// Person interface
type Person interface {
	printBasicInfo()
	printAllInfo()
}

// Student structure
type Student struct {
	name string
	id   int
}

// Teacher structure
type Teacher struct {
	name   string
	branch string
}

func (s Student) printBasicInfo() {
	fmt.Println("Basic Info")
	fmt.Println(s.name)
	fmt.Println()
}

func (s Student) printAllInfo() {
	fmt.Println("All Info")
	fmt.Println(s.name)
	fmt.Println(s.id)
	fmt.Println()
}

func (t Teacher) printBasicInfo() {
	fmt.Println("Basic Info")
	fmt.Println(t.name)
	fmt.Println()
}

func (t Teacher) printAllInfo() {
	fmt.Println("All Info")
	fmt.Println(t.name)
	fmt.Println(t.branch)
	fmt.Println()
}

func introduce(p Person) {
	p.printBasicInfo()
	p.printAllInfo()
}

func main() {
	student1 := Student{"john", 123}
	student2 := Student{"melissa", 456}

	teacher1 := Teacher{"emily", "english"}
	teacher2 := Teacher{"su", "music"}

	introduce(student1)
	introduce(student2)
	introduce(teacher1)
	introduce(teacher2)
}
