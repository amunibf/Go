package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("The square root of 4 is ", math.Sqrt(4))
	fmt.Println("A number from 1-100", rand.Intn(100))

}

func main() {
	foo()
}
