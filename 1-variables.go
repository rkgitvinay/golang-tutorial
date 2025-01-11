package main

import (
	"fmt"
	"sort"
)

const LoginToken string = "login-token" // public variable starts with Capital letter

func main() {
	var username string = "Vinay"

	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.322332
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values
	var defaultVariable int
	fmt.Println(defaultVariable)
	fmt.Printf("Variable is of type: %T \n", defaultVariable)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

	// Array
	fmt.Println("=========Starting Array=========")
	var fruits [3]string
	fruits[0] = "Banana"
	fmt.Println("Fruits", fruits)

	var vegitables = [3]string{"Potato", "Beans", "Onion"}
	fmt.Println("Vegitables", vegitables)

	// Slices
	fmt.Println("=========Starting Slices=========")
	var fruitList = []string{"hello", "slice"}
	fmt.Println("Vegitables", fruitList)

	fruitList = append(fruitList, "append")
	fmt.Println("Vegitables", fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 334
	highScore[2] = 434
	highScore[3] = 534
	// highScore[4] = 44 // throw an error

	highScore = append(highScore, 44) // it will work due to append

	fmt.Println("highScore", highScore)

	sort.Ints(highScore)
	fmt.Println("highScore", highScore)

	// Maps
	fmt.Println("=========Starting Maps=========")
	scores := make(map[string]int) // string is key, and int is values
	scores["Vinay"] = 1
	scores["Ramesh"] = 4
	scores["Raja"] = 5
	fmt.Println("scores", scores)

	delete(scores, "Raja")
	fmt.Println("scores", scores)

	// loop in Maps
	for key, value := range scores {
		fmt.Printf("key %v, value %v \n", key, value)
	}

	for _, value := range scores {
		fmt.Printf("value %v \n", value)
	}

	fmt.Println("========= Starting Structs =========")
	// No inheritance in golang, No Super, parent or child thing in go
	vinay := User{"Vinay", "vinay@gmail.com", true, 25}
	fmt.Println(vinay)
	fmt.Printf("Vinay details are: %+v\n", vinay)
	fmt.Printf("Vinay details are: Name: %v, Email: %v\n", vinay.Name, vinay.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
