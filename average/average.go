// average calculates the average of several numbers.
package main

import (
	"fmt"
)

func main() {
	numbers := [5]float64{71.8, 56.2, 89.5, 98.6, 80.0}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
