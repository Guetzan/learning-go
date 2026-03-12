package main

import "fmt"

func main() {
	//basics
	var pointer1 *int32 = new(int32)
	fmt.Printf("Endereço: %v Valor: %v", pointer1, *pointer1)

	*pointer1 = 20
	fmt.Printf("\nEndereço: %v Valor: %v\n\n", pointer1, *pointer1)

	var year int16 = 2003
	var pointer2 *int16 = new(int16)

	fmt.Printf("Endereço: %v Valor: %v", pointer2, *pointer2)
	fmt.Printf("\nValue of year: %v", year)

	pointer2 = &year
	*pointer2 = 2002
	fmt.Printf("\nEndereço: %v Valor: %v", pointer2, *pointer2)
	fmt.Printf("\nValue of year: %v\n\n", year)

	//using pointers with functions
	var arr [5]uint8 = [5]uint8 {1,2,3,4,5}
	fmt.Println(arr)
	var result [5]uint8 = squareUp(&arr)
	fmt.Println(result)
}

func squareUp(arr *[5]uint8) [5]uint8 {
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}

	return *arr
}

// func squareUp(arr *[5]uint8) {
// 	for i := range *arr {
// 		arr[i] = arr[i] * arr[i]
// 	}
// }