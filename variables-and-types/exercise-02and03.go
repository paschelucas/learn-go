package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func exercise02and03() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println("Exercise 02 and 03 => ", s)
}
