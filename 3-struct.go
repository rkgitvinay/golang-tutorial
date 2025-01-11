package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Creating a struct instance
	p1 := Person{Name: "Vinay", Age: 30}
	fmt.Println("Person 1", p1)

	// Positional syntax
	p2 := Person{"Vinay Singh", 32}
	fmt.Println("Person 2", p2)

	//Accessing and Modifying Fields
	p3 := Person{Name: "Nidhi Singh", Age: 30}
	fmt.Println("Name of p3", p3.Name)
	fmt.Println("Age of p3", p3.Age)

	p3.Name = "Nidhi Vinay Singh"
	p3.Age = 29
	fmt.Println("Name of p3 after update", p3.Name)
	fmt.Println("Age of p3 after update", p3.Age)

	//Structs with Pointers
	p4 := &Person{Name: "Charlie", Age: 22}
	fmt.Println("Name p4", p4.Name)
}
