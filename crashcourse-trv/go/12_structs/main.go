package main

import (
	"fmt"
	"strconv"
)

//define person struct
type Person struct {
	/* firstName string
	lastName  string
	city      string
	gender    string */
	firstName, lastName, city, gender string
	age                               int
}

//greeting method (val receiver)
func (p Person) greet() string {
	return "ahello my name is " + p.firstName + " " + p.lastName + ". I am " +
		strconv.Itoa(p.age)
}

//hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried pointr receiver
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//init person using struct
	person1 := Person{firstName: "Ahmad", lastName: "Munib",
		city: "Gresik", gender: "m", age: 31}
	person3 := Person{firstName: "Samantha", lastName: "Smith",
		city: "Boston", gender: "m", age: 32}

	//alternative
	person2 := Person{"Elias", "Alhanan",
		"Gresik", "m", 3}

	person1.age++
	fmt.Println(person1)
	fmt.Println(person1.greet())
	fmt.Println(person2)
	person2.hasBirthday()
	person2.hasBirthday()
	fmt.Println(person2.greet())

	fmt.Println(person2.lastName)

	person3.getMarried("Williams")
	fmt.Println(person3.greet())

}
