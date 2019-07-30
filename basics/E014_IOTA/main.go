package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d
		e
	)

	const (
		f = iota
		_
		g
		_
		h
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
