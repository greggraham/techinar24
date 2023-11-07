package main

import "fmt"

// YOUR CODE HERE: Define a printLines function.
func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {
	daysOfWeek := [7]string{"Sunday", "Monday", "Tuesday",
		"Wednesday", "Thursday", "Friday", "Saturday"}
	// YOUR CODE HERE: Get a slice containing just the
	// weekdays.
	// Pass that slice to printLines.
	weekDays := daysOfWeek[1:6]
	printLines(weekDays)
}
