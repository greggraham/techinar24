package main

import (
	"fmt"
)

func main() {
	var authors map[string]string = map[string]string{"The Lord of the Rings": "J. R. R. Tolkien", "Mere Christianity": "C. S. Lewis"}

	for title, author := range authors {
		fmt.Printf("%25s : %s\n", title, author)
	}

	fmt.Println("----------------------------------------")

	books := make(map[string]int)
	books["Mere Christianity"]++
	books["The Lord of the Rings"]++
	books["Great Expectations"]++
	books["Mere Christianity"]++
	for title, count := range books {
		fmt.Printf("%-25s : %d\n", title, count)
	}

	fmt.Println("----------------------------------------")

	macros := map[string]string{"chicken": "protein", "avocado": "fat", "apple": "carb", "rice": "carb"}
	for food, macro := range macros {
		fmt.Printf("%-20s : %s\n", food, macro)
	}
}
