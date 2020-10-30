package main

import (
	"fmt"
	"time"
)

func greeting(c chan string) {
	fmt.Println("Hello " + <-c)
	fmt.Println("Welcome " + <-c)
}

func squares(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i * i // send channel
	}

	close(c) // close channel
}

func main() {

	fmt.Println("Main() Start")

	c := make(chan int)

	go squares(c)

	for {
		value, ok := <-c // read channel

		if ok != true {
			fmt.Println("For loop broke !!!")
			break
		}

		fmt.Println(value)
	}

	fmt.Println("Main() End")
	time.Sleep(time.Second)

}
