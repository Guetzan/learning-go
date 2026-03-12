package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = []rune("Clóvis")
	var indexed = name[2]

	fmt.Printf("%v, %T\n", indexed, indexed)

	for index, value := range name {
		fmt.Printf("%v: %v\n", index, value)
	}

	var country = []string{"B", "r", "a", "z", "i", "l"}
	var builder strings.Builder

	for i := range country {
		builder.WriteString(country[i])
	}

	var concatenated string = builder.String()

	fmt.Println(concatenated)
}