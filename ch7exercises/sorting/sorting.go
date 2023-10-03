package main

import (
	"fmt"
	"log"
	"sort"
	"techinar/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	// YOUR CODE HERE: Build a slice containing all the
	// key strings from the "counts" map.
	// Call the sort.Strings method on your slice.
	// Loop through the names in the sorted slice, and
	// print the name and the corresponding count from
	// the map.
	var names []string
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%-20s : %3d\n", name, counts[name])
	}
}
