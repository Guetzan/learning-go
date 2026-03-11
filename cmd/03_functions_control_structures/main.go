package main

import (
	"errors"
	"fmt"
)

func main() {
	var value string = "Hello, World!"
	printMe(value)

	var numerator int = 25
	var denominator int = 4
	var result, remainder, err = intDivision(numerator, denominator)

	switch {
		case err != nil:
			fmt.Println(err.Error())
		case remainder == 0:
			fmt.Printf("The division results in %v", result)
		default:
			fmt.Printf("The division results in %v with a remainder of %v", result, remainder)
	}

	switch remainder {
		case 0:
			fmt.Printf("The division was exact")
		case 1, 2:
			fmt.Printf("The division was close")
		default:
			fmt.Printf("The division was not close")
	}
}

func printMe(value string) {
	fmt.Println(value)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("The division by zero is an illegal operation")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
