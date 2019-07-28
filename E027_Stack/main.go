/**
 * Stack implementation
 */
package main

import "fmt"

// Stack Structure
type Stack struct {
	stack    []int
	top      int
	capacity int
}

func (s *Stack) init(initialCapacity int) {
	s.stack = make([]int, initialCapacity)
	s.capacity = initialCapacity
	s.top = 0
}

func (s *Stack) isFull() bool {
	return s.top == s.capacity
}

func (s *Stack) isEmpty() bool {
	return s.top == 0
}

func (s *Stack) push(value int) bool {
	if !s.isFull() {
		s.stack[s.top] = value
		s.top++
		return true
	}

	return false
}

func (s *Stack) pop() int {
	if !s.isEmpty() {
		s.top--
		tmp := s.stack[s.top]
		return tmp
	}
	return -1
}

func (s *Stack) print() {
	if s.top == 0 {
		fmt.Print("Empty stack")
	}

	for i := 0; i < s.top; i++ {
		fmt.Printf("%d\t", s.stack[i])
	}

	fmt.Println()
}

func main() {
	stack := Stack{}
	stack.init(10)

	stack.print()

	stack.push(1)
	stack.push(2)
	stack.push(3)

	stack.print()

	tmp := stack.pop()
	fmt.Println(tmp)
	stack.print()
}
