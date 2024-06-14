package main

import "fmt"

func main() {
	x := 0

	// if
	if x == 0 {
		fmt.Println("X vale zero")
	} else if x == 1 {
		fmt.Println("X vale 1")
	} else {
		fmt.Println("X nao vale nem 0 nem 1")
	}

	// switch case
	y := 5

	switch {
	case y < 5:
		fmt.Println("Y é menor que 5")
		break
	case y > 5:
		fmt.Println("Y é maior que 5")
		break
	default:
		fmt.Println("Y vale 5")
	}
}