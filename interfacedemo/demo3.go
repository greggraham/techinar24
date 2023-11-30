package main

import "fmt"

func demo3() {
	myList := []interface{}{3, "Bang!", 72, 16, "Pop!"}

	for _, i := range myList {
		number, ok := i.(int)
		if ok {
			fmt.Println(number * 3)
		}
		word, ok := i.(string)
		if ok {
			fmt.Println(word, word, word)
		}
	}
}
