package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "munib"
	c <- "thanks"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
