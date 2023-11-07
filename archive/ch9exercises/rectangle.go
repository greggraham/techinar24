package main

import "fmt"

type rectangle struct {
	length float64
	width  float64
}

func (r *rectangle) info() {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

func (r *rectangle) makeSquare() {
	if r.length > r.width {
		r.length = r.width
	} else {
		r.width = r.length
	}
}

func main() {
	var r rectangle
	r.length = 4.2
	r.width = 2.3
	r.info()
	r.makeSquare()
	r.info()
}
