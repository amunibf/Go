package main

import "fmt"

// var name string = "Munib"

// name := "Munib"

func main() {
	//Using var
	// var name string = "Munib"
	// var name = "Munib"
	// var age int = 37
	var age int32 = 37
	var size float32 = 4.7
	// const isCool = true
	var isCool = true

	//shorthand
	/* name := "Munib"
	email := "munib@gmail.com" */
	name, email := "Munib", "munib@gmail.com"
	// size := 1.3
	isCool = false
	fmt.Println(name, email, age)
	fmt.Printf("%T\n", isCool, size)
}
