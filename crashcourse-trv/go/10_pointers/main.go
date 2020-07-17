package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	//use * to read value from address
	fmt.Println(*b)
	fmt.Println(*&a)

	//change val with pointer
	*b = 100
	fmt.Println(a)

}
