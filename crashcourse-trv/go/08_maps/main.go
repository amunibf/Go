package main

import (
	"fmt"
)

func main() {
	//define map
	emails := make(map[string]string)

	//assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Adi"] = "adi@gmail.com"
	emails["Jone"] = "jone@gmail.com"

	//declare and assign kv
	emailq := map[string]string{"Munib": "munub@gmail.com", "Iis": "iis@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Adi"])

	//delete from map
	delete(emails, "Adi")
	fmt.Println(emails)
	fmt.Println(emailq)

}
