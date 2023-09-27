package main

import (
	"fmt"
	"math"

	"techinar/keyboard"
)

// Read in a positive float64 number and handle all errors.
func getFloatPositive(prompt string) float64 {
	for {
		fmt.Print(prompt)
		number, err := keyboard.GetFloat()
		if number <= 0 || err != nil {
			fmt.Println("Please enter a positive, floating point number.")
		} else {
			return number
		}
	}
}

// Input the size and price for the given pizza.
func getPizzaValues(pizzaNumber int) (float64, float64) {
	prompt := fmt.Sprintf("Enter the diameter for pizza %d (inches): ", pizzaNumber)
	diameter := getFloatPositive(prompt)
	prompt = fmt.Sprintf("Enter the price for pizza %d (dollars): ", pizzaNumber)
	price := getFloatPositive(prompt)
	return diameter, price
}

// Calculate the price per square inch of a pizza.
func calcPSI(pizzaNumber int) float64 {
	diameter, price := getPizzaValues(pizzaNumber)
	area := math.Pow(diameter/2, 2) * math.Pi
	psi := price / area
	fmt.Printf("The price per square inch of pizza %d is $%.2f.\n", pizzaNumber, psi)
	return psi
}

func main() {

	psi1 := calcPSI(1)
	psi2 := calcPSI(2)

	if psi1 < psi2 {
		fmt.Println("Pizza 1 is the better value.")
	} else if psi1 > psi2 {
		fmt.Println("Pizza 2 is the better value.")
	} else {
		fmt.Println("The two pizzas are of equal value.")
	}
}
