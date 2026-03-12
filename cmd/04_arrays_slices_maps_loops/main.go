package main

import "fmt"

func main() {
	//arrays
	fmt.Println("Arrays:")
	var arr [3]int32 = [3]int32{1,2,3}
	fmt.Println(arr[0])
	fmt.Println(arr[1:3])

	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	//slices
	fmt.Println("\nSlices:")
	var slice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with a capacity of %v\n", len(slice), cap(slice))

	slice = append(slice, 7)
	fmt.Printf("The length is %v with a capacity of %v\n", len(slice), cap(slice))

	var slice2 []int32 = []int32{8,9}

	slice = append(slice, slice2...)
	fmt.Println(slice)

	var slice3 []int32 = make([]int32, 3, 8)
	fmt.Println(slice3)

	//maps
	fmt.Println("\nMaps:")
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]int32 {"Ronaldo": 2002, "David": 2003}
	fmt.Println(myMap2)
	var year, ok = myMap2["Jonathan"];
	
	if ok {
		fmt.Println(year)
	} else {
		fmt.Println("The name was not found")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for index, value := range arr {
		fmt.Printf("%v: %v\n", index, value)
	}

	for i := 0; i <= 5; i++ {
		fmt.Printf("%v ", i)
	}
}