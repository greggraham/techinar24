package main

import (
	"fmt"
)

// YOUR CODE HERE: Declare a "rectangle" struct type.
type rectangle struct {
	length float64
	width  float64
}

// YOUR CODE HERE: Define a rectangleInfo function.
func rectangleInfo(rect rectangle) {
	fmt.Println("Length:", rect.length)
	fmt.Println("Width:", rect.width)
}

func makeSquare(rect *rectangle) {
	if rect.length > rect.width {
		rect.length = rect.width
	} else {
		rect.width = rect.length
	}
}

func demo(name string, r *rectangle) {
	fmt.Println("Original size of", name)
	rectangleInfo(*r)
	makeSquare(r)
	fmt.Println("New size of", name)
	rectangleInfo(*r)
	fmt.Println()
}

func main() {
	// YOUR CODE HERE: Create a new rectangle, and set its
	// length and width fields. Pass it to rectangleInfo.
	var r1, r2, r3 rectangle

	r1.length = 12.5
	r1.width = 8.2
	demo("r1", &r1)

	r2.length = 2.2
	r2.width = 15.2
	demo("r2", &r2)

	r3.length = 25.2
	r3.width = 25.2
	demo("r3", &r3)

	fmt.Println("Recheck of r1")
	rectangleInfo(r1)
}
