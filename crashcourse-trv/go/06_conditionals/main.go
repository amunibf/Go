package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 10

	//if else
	if x <= y {
		fmt.Printf("%d is less or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if
	color := "red"

	/* if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is NOT red OR blue")
	} */

	//switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is NOT red OR blue")

	}
}

//&& , ||
