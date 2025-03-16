package hello

import "fmt"

var HelloMessage = "Hello World from Google"

func Hello() {
	fmt.Println(HelloMessage)
	otherFunction()
}

func otherFunction() {
}
