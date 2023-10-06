// First draft of Pizza Price Program with very little modularity.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func ppp() {

	// Input values for pizza 1

	fmt.Print("Enter the diameter for pizza 1 (inches): ")
	diameter1, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Enter the price for pizza 1 (dollars): ")
	price1, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}

	// Input values for pizza 2

	fmt.Print("Enter the diameter for pizza 2 (inches): ")
	diameter2, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Enter the price for pizza 2 (dollars): ")
	price2, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}

	// Pizza 1 calculations

	area1 := math.Pow(diameter1/2, 2) * math.Pi
	psi1 := price1 / area1
	fmt.Printf("The price per square inch of pizza 1 is $%5.2f.\n", psi1)

	// Pizza 2 calculations

	area2 := math.Pow(diameter2/2, 2) * math.Pi
	psi2 := price2 / area2
	fmt.Printf("The price per square inch of pizza 2 is $%5.2f.\n", psi2)

	// Compare values

	if psi1 < psi2 {
		fmt.Println("Pizza 1 is the better value.")
	} else if psi1 > psi2 {
		fmt.Println("Pizza 2 is the better value.")
	} else {
		fmt.Println("The two pizzas are of equal value.")
	}
}
