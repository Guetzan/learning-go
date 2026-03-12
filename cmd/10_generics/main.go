package main

import (
	"errors"
	"fmt"
)

func main() {
	var slice = []int32{}
	var floatSlice = []float32{20.20, 20.0}

	var result, err1 = sum(slice)
	var result2, err2 = sum(floatSlice)

	if err1 != nil {
		fmt.Println(err1.Error())
	} else {
		fmt.Println(result)
	}

	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		fmt.Println(result2)
	}
}

func sum[T int32 | float32 | float64](slice []T) (T, error) {
	var err error
	var sum T = 0

	if isEmpty(slice) {
		err = errors.New("The slice has no values to sum")
		return sum, err
	}

	for _, value := range slice {
		sum += value
	}

	return sum, err
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}