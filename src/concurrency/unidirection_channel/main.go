package main

import (
	"fmt"
	"time"
)

// greet function with receive-only channel param
func greet(name <-chan string) {
	fmt.Println("Hello ", <-name)
}

func main() {

	name := make(chan string) // bi-direction channel
	go greet(name)            // name will be automtically convert to receive-only channel when passed to greet function
	name <- "Mickey"
	// fmt.Println("Main() stopped")
	time.Sleep(time.Second)
}
