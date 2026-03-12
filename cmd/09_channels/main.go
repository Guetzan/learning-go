package main

import "fmt"

func main() {
	var channel = make(chan int)
	go update(channel)

	for i := range channel {
		fmt.Println(i)
	}
}

func update(channel chan int) {
	defer close(channel)

	for i := range 5 {
		channel <- i
	}
}