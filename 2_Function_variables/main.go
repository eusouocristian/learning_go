package main

import (
	"fmt"
	"main/hello"
	"math"
)

// Global variables
// lowercase = Won't be exported to be used from other file
var myNumber float64 = 1.23

// Uppercase = Will exported to be used from other file
var OtherNumber float32 = 2.34

func main() {
	roundedUp := math.Ceil(myNumber)
	roundedDown := math.Floor(myNumber)
	fmt.Println("Original Number:", myNumber)
	fmt.Println("Rounded Up:", roundedUp, "\nRounded down:", roundedDown)

	fmt.Println(hello.HelloMessage)
	hello.Hello()
}
