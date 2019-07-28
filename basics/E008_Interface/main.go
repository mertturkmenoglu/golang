package main

import "fmt"

type Animal interface {
	eat()
	makeSound()
}

func printAnimal(a Animal) {
	fmt.Println(a)
}

type squirrel struct {
	sound string
	color string
}

type dog struct {
	sound string
	name  string
}

func (s squirrel) eat() {
	fmt.Println(s.color, "color squirrel is eating now")
}

func (s squirrel) makeSound() {
	fmt.Println(s.sound)
}

func (d dog) eat() {
	fmt.Println(d.name, "is eating now")
}

func (d dog) makeSound() {
	fmt.Println(d.sound)
}

func main() {
	alf := squirrel{
		sound: "vik",
		color: "brown",
	}

	doggo := dog{
		sound: "hav",
		name:  "GoodBoy",
	}

	alf.eat()
	alf.makeSound()

	doggo.eat()
	doggo.makeSound()

	printAnimal(alf)
	printAnimal(doggo)
}
