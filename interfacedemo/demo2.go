package main

import (
	"fmt"
	"techinar/interfaces"
)

type Dog string

func (d Dog) Speak() {
	fmt.Println(d, "says Bow wow!")
}

func (d Dog) Move() {
	fmt.Println(d, "is running...")
}

type Cat string

func (c Cat) Speak() {
	fmt.Println(c, "says Meow!")
}

func (c Cat) Move() {
	fmt.Println(c, "is prowling...")
}

type Human string

func (h Human) Speak() {
	fmt.Println(h, "says Good morning!")
}

func (h Human) Move() {
	fmt.Println(h, "is walking...")
}

type Cow string

func (c Cow) Speak() {
	fmt.Println(c, "says Moo!")
}

func (c Cow) Move() {
	fmt.Println(c, "is sauntering...")
}

func demo2() {
	animals := []interfaces.Animal{Dog("Fido"), Cat("Morris"), Human("Bob"), Cow("Bessie")}

	for _, a := range animals {
		a.Speak()
		a.Move()
	}
}
