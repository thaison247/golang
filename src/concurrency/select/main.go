package main

import (
	"fmt"
	"time"
)

func service1(c chan string) {
	time.Sleep(4 * time.Second)
	c <- "Service 1 channel content"
}

func service2(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Service 2 channel content"
}

func main() {

	// create channels
	chan1 := make(chan string)
	chan2 := make(chan string)
	chan3 := make(chan string, 2)
	chan4 := make(chan string, 2)

	// run goroutines
	go service1(chan1)
	go service2(chan2)

	// send data to chan3 & chan4
	chan3 <- "Hello"
	chan3 <- "Konichiwa"
	chan4 <- "Xin Chao"
	chan4 <- "Bonjour"

	select {
	case res := <-chan1: // blocking case
		fmt.Println("Result in first case: ", res)
	case res := <-chan2: // blocking case
		fmt.Println("Result in sencond case: ", res)
	case res := <-chan3: // unblocking case
		fmt.Println("Result in third case: ", res)
	case res := <-chan4: // unblocking case
		fmt.Println("Result in fouth case: ", res)
	}

	fmt.Println("Main() ended")
}
