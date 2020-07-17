package main

import {
	"fmt"
	"errors"
	"math"
}

/* func main() {
	fmt.Println("Hello world!")

	var x int = 5
	var y int = 7
	var sum int = x + y
	fmt.Println(sum)

	n := -5
	if n > 2 {
		fmt.Println("more than 2")
	} else if n < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("btween 0 and 2")
	}

	var a [5]int
	b := []int{8, 4, 7, 4, 1, 7, 4}
	a[2] = 4
	b = append(b, 89)
	fmt.Println(a)
	fmt.Println(b)

	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 5
	delete(vertices, "square")
	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

	for i := 0; i < 6; i++ {
		fmt.Print(i)
	}

	z := 2
	for z < 6 {
		fmt.Println(z)
		z++
	}

	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	m := make(map[string]string)
	m["a"] = "aplha"
	m["b"] = "beta"
	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}
} */

/* func main() {
	result := sum(2, 3)
	fmt.Println(result)
} */


/* func sum(x int, y int) int {
	return x + y
} */

func main() {
	result := sum(2, 3)
	fmt.Println(result)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undef for negatif number")
	}
	return math.Sqrt(x) 
}
