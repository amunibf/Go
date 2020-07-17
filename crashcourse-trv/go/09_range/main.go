package main

import (
	"fmt"
)

func main() {
	ids := []int{23, 54, 32, 55, 12, 87}

	//loop thru ids
	for i, id := range ids {
		fmt.Printf("%d - ID : %d\n", i, id)
	}

	//not using index
	for _, id := range ids {
		fmt.Printf("ID : %d\n", id)
	}

	//add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum : ", sum)

	//Range with map
	emails := map[string]string{"Munib": "munub@gmail.com", "Iis": "iis@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}

	//get key only
	for k := range emails {
		fmt.Println("Name: " + k)
	}

	//get value only
	for _, v := range emails {
		fmt.Println("Email: " + v)
	}
}
