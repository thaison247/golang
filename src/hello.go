package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func greeting(c chan string) {
	fmt.Println("Hello " + <-c)
	fmt.Println("Welcome " + <-c)
}

func main() {

	fmt.Println("Start")

	c := make(chan string)

	go greeting(c)

	c <- "Mera"

	close(c)

	fmt.Println("End")
	time.Sleep(time.Second)

}
