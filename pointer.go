package main

import "fmt"

func incrementByValue(val int) {
	val += 1
}

func incrementByReference(ptrval *int) {
	*ptrval += 1
}

func main() {
	num := 5
	incrementByValue(5)
	fmt.Println("After incrementByValue:", num)

	incrementByReference(&num)
	fmt.Println("After incrementByReference:", num)
}
