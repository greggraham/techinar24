package main

import (
	"os"
	"strconv"
)

func main() {
	var d int
	var err error
	if len(os.Args) < 2 {
		d = 0
	} else {
		d, err = strconv.Atoi(os.Args[1])
		if err != nil {
			d = 0
		}
	}
	switch d {
	case 1:
		demo1()
	case 2:
		demo2()
	case 3:
		demo3()
	default:
		demo1()
		demo2()
		demo3()
	}
}
