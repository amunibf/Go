package main

import (
	"fmt"
)

func main() {
	//Arrays
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Peach"

	//declare and assign array
	buahArr := [2]string{"Jambu", "Mangga"}

	//slice
	buahSlice := []string{"Kedondong", "Manggis", "Pear", "Strawberry"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(buahArr)
	fmt.Println(buahArr[1])
	fmt.Println(buahSlice)
	fmt.Println(len(buahSlice))
	fmt.Println(buahSlice[0:3])
}
