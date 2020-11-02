package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func sender(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	close(c)
}

func main() {
	fmt.Println("Started")

	c := make(chan int, 3)

	go sender(c)

	for v := range c {

		val := v

		length := len(c)

		fmt.Printf("Channel c's length after read value %d is %d\n", val, length)
	}

	// c <- 5

	fmt.Println("Ended")
}
