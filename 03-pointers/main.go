package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers section")

	// var ptr *int
	// fmt.Println("Value of ptr is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of ptr is ", ptr)
	fmt.Println("Value of ptr is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is", myNumber)

}
