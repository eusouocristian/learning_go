package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// Create a var and assign a value
	var string1 = "First String"
	string2 := "Second String"

	// Create a var without declaring a value
	// Requires declaring the type
	var number1 uint16 = 11
	number2 := 2.3

	//Array Declaration:
	var bookings = [50]string{"Nana", "Nicole", "Cristian"}
	for i := 0; i < len(bookings); i++ {
		fmt.Println(bookings[i])
		if bookings[i] == "" {
			break
		}
	}

	fmt.Printf("Number 1 is %d and number 2 is %.1f\n", number1, number2)
	fmt.Printf("String 1 is %s and String 2 is %s", string1, string2)
}
