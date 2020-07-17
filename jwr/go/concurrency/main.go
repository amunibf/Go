package main

import (
	"fmt"
	"time"
)

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// fmt.Println(i, thing)
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}

func main() {
	c := make(chan string)
	go count("munib", c)
	for msg := range c {
		// msg, open := <-c
		/* if !open {
			break
		} */
		fmt.Println(msg)
	}

	/* var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("munib")
		wg.Done()
	}()

	wg.Wait() */

	// go count("munib")

	// go count("iis")
	// time.Sleep(time.Millisecond * 500)
	// fmt.Scanln()

}
