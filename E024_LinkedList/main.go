/**
 * Linear linked list example
 */
package main

import "fmt"

type Node struct {
	value int
	next *Node
}

func main() {
	var head *Node
	head = nil
	var tmp *Node

	fmt.Println("Add to empty linked list")
	head = &Node{value: 1, next: nil}

	fmt.Println("Add before head")
	tmp = &Node{value: -1, next: head}
	head = tmp

	fmt.Println("Add between nodes")
	mid := Node{value: 0, next: head.next}
	head.next = &mid

	tmp = head
	fmt.Println(tmp.value)
	
	tmp = head.next
	fmt.Println(tmp.value)
	
	tmp = tmp.next
	fmt.Println(tmp.value)
}