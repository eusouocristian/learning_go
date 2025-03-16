package main

import (
	"fmt"
	"main/welcome"
	"math"

	"golang.org/x/example/hello/reverse"
)

// variable here is availabe for all functions within this file
var myNumber = 1.23

func main() {
	fmt.Println("Hello Go Language")
	fmt.Println(math.Ceil(myNumber))  //rounded up
	fmt.Println(math.Round(myNumber)) //rounded down

	fmt.Println(welcome.Japanese)

	// Have to install the package
	// go get golang.org/x/example/hello/reverse
	fmt.Println(reverse.String("Mynameiscristian"))
}
