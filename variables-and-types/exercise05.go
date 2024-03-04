package main

import "fmt"

type myType05 int

var myVar05 myType05

func exercise05() {
	myVar05 = 42
	myVarConverted := int(myVar05)

	fmt.Print("Exercise 05 => ")
	fmt.Printf("%T\n", myVarConverted)
}
